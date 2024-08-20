package chinaz

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/ent"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/ent/icp"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/persistent/db"
	"go.uber.org/zap"
)

var _ application.Application = (*Chinaz)(nil)

type Chinaz struct {
	*application.EmptyApplication
	config config
	db     *db.Client
}

func New() *Chinaz {
	return &Chinaz{
		EmptyApplication: application.NewEmptyApplication(),
		config:           config{},
	}
}

func (c *Chinaz) SetDB(db *db.Client) {
	c.db = db
}

func (c *Chinaz) Configuration() configuration.Configuration {
	return &c.config
}

func (c *Chinaz) Run(ctx context.Context) {
}

func (c *Chinaz) Query(ctx context.Context, _domain string) (*ent.Icp, bool, error) {
	domain, ok := to_valid_domain(_domain)
	if !ok {
		return nil, false, fmt.Errorf("invalid domain: %s", _domain)
	}
	fmt.Println(domain)
	res, err := c.db.Icp.Query().Where(icp.Host(domain)).First(ctx)
	if err == nil {
		return res, true, nil
	} else if !ent.IsNotFound(err) {
		return nil, false, err
	}
	c.Logger.Debug("query from chinaz", zap.String("domain", domain))
	resp, err := c.query(domain)
	if err != nil {
		return nil, false, err
	}

	res, err = c.db.Icp.Create().
		SetHost(domain).
		SetCompany(resp.Result.CompanyName).
		SetOwner(resp.Result.Owner).
		SetType(resp.Result.CompanyType).
		SetHomepage(resp.Result.MainPage).
		SetPermit(resp.Result.SiteLicense).
		Save(ctx)
	return res, false, err
}

type IcpData struct {
	Reason    string `json:"Reason"`
	StateCode int    `json:"StateCode"`
	Result    struct {
		CompanyName string `json:"CompanyName"`
		CompanyType string `json:"CompanyType"`
		MainPage    string `json:"MainPage"`
		Owner       string `json:"Owner"`
		SiteLicense string `json:"SiteLicense"`
		SiteName    string `json:"SiteName"`
		VerifyTime  string `json:"VerifyTime"`
	} `json:"Result"`
}

func (c *Chinaz) query(domain string) (*IcpData, error) {
	url, _ := url.Parse("https://openapi.chinaz.net/v1/1001/icp")
	q := url.Query()
	q.Set("domain", domain)
	q.Set("APIKey", c.config.Token)
	q.Set("ChinazVer", "1.0")
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result *IcpData
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.Result.CompanyType == "" {
		result.Result.CompanyType = "INVALID"
	}

	return result, nil
}

// func (*Chinaz) sign(data string) map[string]string {
// 	parted := strings.Split(data, ".")
// 	// rand from 100 to 999
// 	rand := 266

// 	res := []string{}

// 	for idx, part := range parted {
// 		ascii_sum := 0
// 		for _, char := range part {
// 			ascii_sum += int(char)
// 		}
// 		if idx == len(parted)-1 {
// 			res = append(res, fmt.Sprintf("%d", ascii_sum+rand))
// 		} else {
// 			res = append(res, fmt.Sprintf("%d", ascii_sum))
// 		}
// 	}

// 	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
// 	key := fmt.Sprintf("%d,%s", rand, strings.Join(res, ","))

// 	h := md5.New()
// 	h.Write([]byte(fmt.Sprintf("%s*#|&%s", key, ts)))

// 	return map[string]string{
// 		"rd":    fmt.Sprintf("%d", rand),
// 		"key":   key,
// 		"ts":    ts,
// 		"token": fmt.Sprintf("%x", h.Sum(nil)),
// 	}
// }

// type LatestIcpData struct {
// 	Message string `json:"msg"`
// 	Code    int    `json:"code"`
// 	Data    struct {
// 		Data struct {
// 			List []struct {
// 				AddrC       string `json:"addrC"`
// 				AddrP       string `json:"addrP"`
// 				Classify    string `json:"classify"`
// 				Company     string `json:"comName"`
// 				ContentType string `json:"contentType"`
// 				Homepage    string `json:"homePage"`
// 				Host        string `json:"host"`
// 				Prtmit      string `json:"prtmit"`
// 				Owner       string `json:"owner"`
// 				Type        string `json:"type"`
// 				UpTime      string `json:"upTime"`
// 				VerifyTime  string `json:"verifyTime"`
// 				WebName     string `json:"webName"`
// 			} `json:"list"`
// 		} `json:"data"`
// 	} `json:"data"`
// }

// func (c *Chinaz) QueryLatestIcpData() (*LatestIcpData, error) {
// 	data := map[string]any{"day": "0", "start": "", "end": "", "typ": "", "suffix": "", "suffixEx": "", "addrp": "", "addrc": "", "webStatus": "", "ex": "1", "keyword": "", "pageNo": 1, "module": "provinces"}
// 	body, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sign := c.sign(string(body))
// 	res, err := http.NewRequest("POST", "https://icp.chinaz.com/provinces/api/queryIcpData", strings.NewReader(string(body)))
// 	if err != nil {
// 		return nil, err
// 	}
// 	res.Header.Set("Content-Type", "application/json")
// 	res.Header.Set("Cookie", c.config.Cookie)
// 	for k, v := range sign {
// 		res.Header.Set(k, v)
// 	}
// 	res.Header.Set("Module", "provinces")
// 	res.Header.Set("Referer", "https://icp.chinaz.com/provinces")
// 	res.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
// 	res.Header.Set("X-Requested-With", "XMLHttpRequest")

// 	client := &http.Client{}
// 	resp, err := client.Do(res)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var result LatestIcpData
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if result.Code != 0 {
// 		return nil, fmt.Errorf("chinaz: %s", result.Message)
// 	}

// 	return &result, nil
// }
