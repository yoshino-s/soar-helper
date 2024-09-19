package chinaz

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/ent"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/ent/icp"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/persistent/db"
	"go.uber.org/zap"
)

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

func (c *Chinaz) BatchQuery(ctx context.Context, domains []string) ([]*ent.Icp, []bool, error) {
	results := make([]*ent.Icp, len(domains))
	valid_domains := make([]string, len(domains))
	valid_domains_set := make(map[string]struct{}, len(domains))
	cached := make([]bool, len(domains))

	for idx, _domain := range domains {
		domain, ok := to_valid_domain(_domain)
		if !ok {
			continue
		}
		valid_domains[idx] = domain
		valid_domains_set[domain] = struct{}{}
	}

	queryRes, err := c.db.Icp.Query().Where(icp.HostIn(common.MapKeys(valid_domains_set)...)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	cachedMap := make(map[string]*ent.Icp)
	fetchedMap := make(map[string]*ent.Icp)

	for _, res := range queryRes {
		cachedMap[res.Host] = res
		delete(valid_domains_set, res.Host)
	}
	for domain := range valid_domains_set {
		if _, ok := fetchedMap[domain]; ok {
			continue
		}
		fetchedMap[domain], err = c.query(ctx, domain)
		if err != nil {
			c.Logger.Error("query from chinaz", zap.String("domain", domain), zap.Error(err))
		}
	}

	for idx, domain := range valid_domains {
		if query := cachedMap[domain]; query != nil {
			results[idx] = query
			cached[idx] = true
		} else {
			results[idx] = fetchedMap[domain]
			cached[idx] = false
		}
	}

	return results, cached, nil
}

func (c *Chinaz) Query(ctx context.Context, _domain string) (*ent.Icp, bool, error) {
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

func (c *Chinaz) query(ctx context.Context, domain string) (*ent.Icp, error) {
	c.Logger.Debug("query from chinaz", zap.String("domain", domain))

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

	res, err := c.db.Icp.Create().
		SetHost(domain).
		SetCompany(result.Result.CompanyName).
		SetOwner(result.Result.Owner).
		SetType(result.Result.CompanyType).
		SetHomepage(result.Result.MainPage).
		SetPermit(result.Result.SiteLicense).
		SetWebName(result.Result.SiteName).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	c.Logger.Debug("query from chinaz", zap.String("domain", domain), zap.Any("result", res), zap.Error(err))

	res, err = c.db.Icp.Query().Where(icp.ID(res.ID)).First(ctx)

	return res, err
}
