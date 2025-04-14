package beian

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/errors"
	"github.com/yoshino-s/soar-helper/internal/ent"
	"github.com/yoshino-s/soar-helper/internal/ent/icp"
	"github.com/yoshino-s/soar-helper/internal/persistent/db"
	"go.uber.org/zap"
)

type Beian struct {
	*application.EmptyApplication
	config config
	db     *db.Client
}

func New() *Beian {
	return &Beian{
		EmptyApplication: application.NewEmptyApplication(),
		config:           config{},
	}
}

func (c *Beian) SetDB(db *db.Client) {
	c.db = db
}

func (c *Beian) Configuration() configuration.Configuration {
	return &c.config
}

func (c *Beian) Run(ctx context.Context) {
}

type resultMapItem struct {
	icp    *ent.Icp
	cached bool
}

func (c *Beian) BatchQuery(ctx context.Context, domains []string) ([]*ent.Icp, []bool, error) {
	results := make([]*ent.Icp, len(domains))
	cached := make([]bool, len(domains))

	resultMap := make(map[string]resultMapItem)

	valid_domains := make([]string, len(domains))
	valid_domains_set := mapset.NewSet[string]()

	for idx, _domain := range domains {
		domain, ok := to_valid_domain(_domain)
		if !ok {
			continue
		}
		valid_domains[idx] = domain
		valid_domains_set.Add(domain)
	}

	queryRes, err := c.db.Icp.Query().Where(icp.HostIn(valid_domains_set.ToSlice()...)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	for _, res := range queryRes {
		resultMap[res.Host] = resultMapItem{icp: res, cached: true}
		valid_domains_set.Remove(res.Host)
	}
	for domain := range valid_domains_set.Iter() {
		if _, ok := resultMap[domain]; ok {
			continue
		}
		res, err := c.query(ctx, domain)
		if err != nil {
			c.Logger.Error("query failed", zap.String("domain", domain), zap.Error(err))
		}
		resultMap[domain] = resultMapItem{icp: res, cached: false}
	}

	for idx, domain := range valid_domains {
		if res, ok := resultMap[domain]; ok && res.icp != nil {
			results[idx] = res.icp
			cached[idx] = res.cached
		} else {
			results[idx] = &ent.Icp{
				Host:      domain,
				Type:      "INVALID",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			cached[idx] = false
		}
	}

	return results, cached, nil
}

func (c *Beian) Query(ctx context.Context, _domain string) (*ent.Icp, bool, error) {
	domain, ok := to_valid_domain(_domain)
	if !ok {
		return nil, false, fmt.Errorf("invalid domain: %s", _domain)
	}
	res, err := c.db.Icp.Query().Where(icp.Host(domain)).First(ctx)
	if err == nil {
		return res, true, nil
	} else if !ent.IsNotFound(err) {
		return nil, false, err
	}

	r, err := c.query(ctx, domain)
	if err != nil {
		return nil, false, err
	}
	return r, false, nil
}

func (c *Beian) query(ctx context.Context, domain string) (*ent.Icp, error) {
	c.Logger.Debug("config", zap.Any("config", c.config))
	var icp *ent.Icp
	var err error
	if c.config.ChinazToken != "" {
		icp, err = c.chinazQuery(ctx, domain)
	} else if c.config.WerplusKey != "" {
		icp, err = c.werplusQuery(ctx, domain)
	} else {
		return nil, errors.New("no beian query service", 500)
	}

	if err != nil {
		return nil, err
	}

	// save icp
	icp, err = c.db.Icp.Create().
		SetIcp(icp).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return c.db.Icp.Get(ctx, icp.ID)
}

type WerplusIcpDataItem struct {
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
	UpdateRecordTime string `json:"updateRecordTime"`
}

type WerplusIcpData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Params struct {
			List []WerplusIcpDataItem `json:"list"`
		} `json:"params"`
	} `json:"data"`
	ExecTime float64 `json:"exec_time"`
	IP       string  `json:"ip"`
}

func (c *Beian) werplusQuery(ctx context.Context, domain string) (*ent.Icp, error) {
	c.Logger.Debug("query from werplus", zap.String("domain", domain))

	url, _ := url.Parse("https://api2.wer.plus/api/offline_domain_icp")
	q := url.Query()
	q.Set("search", domain)
	q.Set("key", c.config.WerplusKey)
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result *WerplusIcpData
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	res := WerplusIcpDataItem{}

	if result.Code != 200 {
		return nil, errors.New(fmt.Sprintf("werplus query failed: %s", result.Msg), 500)
	}

	if result.Code == 404 || len(result.Data.Params.List) == 0 {
		res.NatureName = "INVALID"
	} else {
		res = result.Data.Params.List[0]
	}

	createdAt := time.Now()
	if res.UpdateRecordTime != "" {
		createdAt, err = time.Parse(time.DateTime, res.UpdateRecordTime)
		if err == nil {
			createdAt = time.Now()
		}
	}

	icp := ent.Icp{
		Host:      domain,
		Company:   res.UnitName,
		Owner:     res.LeaderName,
		Type:      res.NatureName,
		Homepage:  "",
		Permit:    res.MainLicence,
		WebName:   "",
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
	}

	c.Logger.Debug("query from werplus", zap.String("domain", domain), zap.Any("icp", icp))

	return &icp, nil
}

type ChinazIcpData struct {
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

func (c *Beian) chinazQuery(ctx context.Context, domain string) (*ent.Icp, error) {
	c.Logger.Debug("query from chinaz", zap.String("domain", domain))

	url, _ := url.Parse("https://openapi.chinaz.net/v1/1001/icp")
	q := url.Query()
	q.Set("domain", domain)
	q.Set("APIKey", c.config.ChinazToken)
	q.Set("ChinazVer", "1.0")
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result *ChinazIcpData
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.Result.CompanyType == "" {
		result.Result.CompanyType = "INVALID"
	}

	icp := ent.Icp{
		Host:      domain,
		Company:   result.Result.CompanyName,
		Owner:     result.Result.Owner,
		Type:      result.Result.CompanyType,
		Homepage:  result.Result.MainPage,
		Permit:    result.Result.SiteLicense,
		WebName:   result.Result.SiteName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.Logger.Debug("query from werplus", zap.String("domain", domain), zap.Any("icp", icp))

	return &icp, nil
}
