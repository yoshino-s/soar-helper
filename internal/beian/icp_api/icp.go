package icp_api

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/yoshino-s/soar-helper/internal/proxy"
	"go.uber.org/zap"
)

type HttpRequestError struct {
	Response *http.Response
}

func (e *HttpRequestError) Error() string {
	return fmt.Sprintf("http request error, status code: %d", e.Response.StatusCode)
}

type IcpQueryData struct {
	Msg    string `json:"msg"`
	Code   int    `json:"code"`
	Params struct {
		Total int `json:"total"`
		List  []struct {
			ContentTypeName  string `json:"contentTypeName"`
			Domain           string `json:"domain"`
			DomainId         int    `json:"domainId"`
			LeaderName       string `json:"leaderName"`
			LimitAccess      string `json:"limitAccess"`
			MainId           int    `json:"mainId"`
			MainLicence      string `json:"mainLicence"`
			NatureName       string `json:"natureName"`
			ServiceId        int    `json:"serviceId"`
			ServiceLicence   string `json:"serviceLicence"`
			UnitName         string `json:"unitName"`
			UpdateRecordTime Time   `json:"updateRecordTime"`
		} `json:"list"`
	} `json:"params"`
}

type IcpApi struct {
	uid         string
	token       string
	tokenExpire time.Time

	proxy *proxy.Proxy

	httpClient *http.Client
}

type proxyKey struct{}

func New(proxy *proxy.Proxy) *IcpApi {
	randByte := make([]byte, 16)
	rand.Read(randByte)

	return &IcpApi{
		uid:   hex.EncodeToString(randByte),
		token: "",
		proxy: proxy,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
			Transport: &http.Transport{
				Proxy: func(req *http.Request) (*url.URL, error) {
					currentProxy, _ := req.Context().Value(proxyKey{}).(*url.URL)
					return currentProxy, nil
				},
			},
		},
	}
}

func (api *IcpApi) RefreshToken(ctx context.Context, force bool) error {
	if time.Now().Before(api.tokenExpire) && !force {
		return nil
	}

	zap.L().Debug("refresh token", zap.String("uid", api.uid))

	timeStampInt := time.Now().UnixMilli()
	timeStampStr := strconv.FormatInt(timeStampInt, 10)

	body := url.Values{}
	body.Set("authKey", fmt.Sprintf("%x", md5.Sum([]byte("testtest"+timeStampStr))))
	body.Set("timeStamp", timeStampStr)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "https://hlwicpfwc.miit.gov.cn/icpproject_query/api/auth", strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Referer", "https://beian.miit.gov.cn/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Cookie", fmt.Sprintf("__jsluid_s=%s", api.uid))

	response, err := api.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return &HttpRequestError{Response: response}
	}

	var resp struct {
		Code   int    `json:"code"`
		Msg    string `json:"msg"`
		Params struct {
			Bussiness string `json:"bussiness"`
			Expire    int    `json:"expire"`
		} `json:"params"`
	}

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("error code: %d, msg: %s", resp.Code, resp.Msg)
	}

	zap.L().Debug("refresh token success", zap.String("token", resp.Params.Bussiness), zap.Duration("expire", time.Duration(resp.Params.Expire)*time.Millisecond))
	api.token = resp.Params.Bussiness
	api.tokenExpire = time.Now().Add(time.Duration(resp.Params.Expire) * time.Millisecond)

	return nil
}

func isUnretriable(err error) bool {
	if err, ok := err.(*HttpRequestError); ok && err.Response.StatusCode != 403 {
		return true
	}

	return false
}

func (api *IcpApi) Query(ctx context.Context, sign string, domain string) (*IcpQueryData, error) {
	blacklist := []string{
		"dbappsecurity.com.cn",
	}
	if slices.Contains(blacklist, domain) {
		return nil, fmt.Errorf("domain %s is in blacklist", domain)
	}

	retry_time := 0
	for {
		retry_time++

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		var res *IcpQueryData
		var err error

		api.proxy.Do(func(url *url.URL) {
			res, err = api.query(context.WithValue(
				ctx,
				proxyKey{},
				url,
			), sign, domain)
		})
		if err == nil {
			return res, nil
		}

		zap.L().Info("query ICP data failed", zap.String("domain", domain), zap.Error(err), zap.Int("retry_time", retry_time))

		if isUnretriable(err) || retry_time >= 5 {
			return nil, err
		}
	}
}

func (api *IcpApi) query(ctx context.Context, sign string, domain string) (*IcpQueryData, error) {
	if err := api.RefreshToken(ctx, false); err != nil {
		return nil, err
	}

	zap.L().Debug("querying ICP data", zap.String("domain", domain))
	body, _ := json.Marshal(map[string]string{
		"unitName":    domain,
		"serviceType": "1",
		"pageSize":    "",
		"pageNum":     "",
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://hlwicpfwc.miit.gov.cn/icpproject_query/api/icpAbbreviateInfo/queryByCondition/", strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Token", api.token)
	req.Header.Set("Sign", sign)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://beian.miit.gov.cn/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Cookie", fmt.Sprintf("__jsluid_s=%s", api.uid))

	response, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, &HttpRequestError{Response: response}
	}

	var resp IcpQueryData
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("error code: %d, msg: %s", resp.Code, resp.Msg)
	}

	zap.L().Debug("query ICP data success", zap.String("domain", domain), zap.Any("data", resp))

	return &resp, nil
}
