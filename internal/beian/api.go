package beian

import (
	"context"
	"sync"
	"time"

	"github.com/go-errors/errors"
	"github.com/hashicorp/go-set/v3"
	"github.com/sourcegraph/conc/iter"
	"github.com/yoshino-s/go-app/telemetry"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/log"
	"github.com/yoshino-s/soar-helper/internal/beian/icp_api"
	"github.com/yoshino-s/soar-helper/internal/ent"
	"github.com/yoshino-s/soar-helper/internal/ent/icp"
	"github.com/yoshino-s/soar-helper/internal/persistent/db"

	"github.com/yoshino-s/soar-helper/internal/proxy"
	"go.uber.org/zap"
)

type Beian struct {
	*application.EmptyApplication
	config config
	db     *db.Client `inject:""`
	icpApi *icp_api.IcpApi

	dbLock sync.Mutex
}

func New() *Beian {
	return &Beian{
		EmptyApplication: application.NewEmptyApplication("beian"),
		config:           config{},
		dbLock:           sync.Mutex{},
	}
}

func (c *Beian) Set(proxy *proxy.Proxy) {
	c.icpApi = icp_api.New(proxy)
}

func (c *Beian) Configuration() configuration.Configuration {
	return &c.config
}

func (c *Beian) Initialize(ctx context.Context) {
	if c.config.MiitSign == "" {
		panic("miit_sign is required")
	}

	c.icpApi.RefreshToken(ctx, true)
}

type resultMapItem struct {
	icp    *ent.Icp
	cached bool
}

func (c *Beian) BatchQuery(ctx context.Context, domains []string, noCache bool) ([]*ent.Icp, []bool, map[string]string, error) {
	results := make([]*ent.Icp, len(domains))
	cached := make([]bool, len(domains))
	errorsMap := make(map[string]string)

	resultMap := make(map[string]resultMapItem)

	valid_domains := make([]string, len(domains))
	valid_domains_set := set.New[string](0)

	for idx, _domain := range domains {
		domain, ok := toValidDomain(_domain)
		if !ok {
			errorsMap[_domain] = "invalid domain"
			continue
		}
		valid_domains[idx] = domain
		valid_domains_set.Insert(domain)
	}

	if !noCache {
		queryRes, err := c.db.Icp.Query().Where(icp.HostIn(valid_domains_set.Slice()...)).All(ctx)
		if err != nil {
			return nil, nil, nil, err
		}

		for _, res := range queryRes {
			resultMap[res.Host] = resultMapItem{icp: res, cached: true}
			valid_domains_set.Remove(res.Host)
		}
	}

	iter.ForEach(valid_domains_set.Slice(), func(_domain *string) {
		domain := *_domain
		if _, ok := resultMap[domain]; ok {
			return
		}
		res, err := c.query(ctx, domain)

		if err != nil {
			c.Logger.Error("query failed", zap.String("domain", domain), zap.Error(err), log.Context(ctx))
			errorsMap[domain] = err.Error()
		}
		resultMap[domain] = resultMapItem{icp: res, cached: false}
	})

	for idx, domain := range valid_domains {
		if res, ok := resultMap[domain]; ok && res.icp != nil {
			results[idx] = res.icp
			cached[idx] = res.cached
		} else {
			results[idx] = &ent.Icp{
				Host:      domain,
				Type:      "ERROR",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			cached[idx] = false
		}
	}

	return results, cached, errorsMap, nil
}

func (c *Beian) Query(ctx context.Context, _domain string, noCache bool) (*ent.Icp, bool, error) {
	domain, ok := toValidDomain(_domain)
	if !ok {
		return nil, false, telemetry.ReportError(ctx, errors.Errorf("invalid domain: %s", _domain))
	}
	if !noCache {
		res, err := c.db.Icp.Query().Where(icp.Host(domain)).First(ctx)
		if err == nil {
			return res, true, nil
		} else if !ent.IsNotFound(err) {
			return nil, false, err
		}
	}

	r, err := c.query(ctx, domain)
	if err != nil {
		return nil, false, err
	}
	return r, false, nil
}

func (c *Beian) query(ctx context.Context, domain string) (*ent.Icp, error) {
	var icp *ent.Icp
	var err error
	icp, err = c.icpQueryQuery(ctx, domain)

	if err != nil {
		return nil, err
	}

	c.Logger.Debug("query result", zap.String("domain", domain), zap.Any("icp", icp), log.Context(ctx))

	c.dbLock.Lock()
	defer c.dbLock.Unlock()

	// save icp
	id, err := c.db.Icp.Create().
		SetHost(icp.Host).
		SetCity(icp.City).
		SetProvince(icp.Province).
		SetCompany(icp.Company).
		SetOwner(icp.Owner).
		SetType(icp.Type).
		SetHomepage(icp.Homepage).
		SetPermit(icp.Permit).
		SetWebName(icp.WebName).
		SetCreatedAt(icp.CreatedAt).
		OnConflict().
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return nil, err
	}

	return c.db.Icp.Get(ctx, id)
}

func (c *Beian) icpQueryQuery(ctx context.Context, domain string) (*ent.Icp, error) {
	c.Logger.Debug("query from icp-query", zap.String("domain", domain), log.Context(ctx))

	result, err := c.icpApi.Query(ctx, c.config.MiitSign, domain)
	if err != nil {
		return nil, err
	}

	if result.Params.Total == 0 || len(result.Params.List) == 0 {
		return &ent.Icp{
			Host:      domain,
			Type:      "INVALID",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	} else {

		return &ent.Icp{
			Host:      domain,
			City:      "",
			Province:  "",
			Company:   result.Params.List[0].UnitName,
			Owner:     result.Params.List[0].LeaderName,
			Type:      result.Params.List[0].NatureName,
			Homepage:  "",
			Permit:    result.Params.List[0].MainLicence,
			WebName:   "",
			CreatedAt: time.Time(result.Params.List[0].UpdateRecordTime),
		}, nil
	}
}
