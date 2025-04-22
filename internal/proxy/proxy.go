package proxy

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	kuaidailigo "github.com/yoshino-s/kuaidaili-go"
	"go.uber.org/zap"
)

type Strategy string

const (
	StrategyRandom     Strategy = "random"
	StrategyRoundRobin Strategy = "roundrobin"
	StrategyFirst      Strategy = "first"
)

var _ application.Application = (*Proxy)(nil)

type Proxy struct {
	*application.EmptyApplication
	client *kuaidailigo.DPSOrderClient

	config config

	proxies []*kuaidailigo.DPSProxy

	rwLock  sync.RWMutex
	current atomic.Int32
}

func New() *Proxy {
	return &Proxy{
		EmptyApplication: application.NewEmptyApplication("proxy"),
		proxies:          make([]*kuaidailigo.DPSProxy, 0),
		rwLock:           sync.RWMutex{},
		current:          atomic.Int32{},
	}
}

func (p *Proxy) Configuration() configuration.Configuration {
	return &p.config
}

func (p *Proxy) Setup(ctx context.Context) {
	p.client = &kuaidailigo.DPSOrderClient{
		OrderClient: kuaidailigo.NewOrderClient(
			p.config.SecretId,
			p.config.SecretKey,
		),
	}
}

func (proxy *Proxy) GetHttp() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxy.GetHttpProxy(),
		},
	}
	return client
}

func (proxy *Proxy) GetHttpProxy() func(*http.Request) (*url.URL, error) {
	return func(r *http.Request) (*url.URL, error) {
		proxyURL := proxy.Next(r.Context())
		if proxyURL == nil {
			return nil, fmt.Errorf("no available proxy")
		}
		r.URL.Scheme = "https"
		zap.L().Info("using proxy", zap.String("proxy", proxyURL.String()))
		return proxyURL, nil
	}
}
