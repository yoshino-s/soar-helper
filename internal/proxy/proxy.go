package proxy

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/bep/debounce"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/log"
	kuaidailigo "github.com/yoshino-s/kuaidaili-go"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var _ application.Application = (*Proxy)(nil)

type Proxy struct {
	*application.EmptyApplication
	config config

	client      *kuaidailigo.AccountClient
	proxyClient *kuaidailigo.TPSOrderClient
	proxyUrl    *url.URL

	proxyMutex sync.Mutex

	debounce func(f func())
}

func New() *Proxy {
	return &Proxy{
		EmptyApplication: application.NewEmptyApplication("proxy"),
		config:           config{},
		proxyMutex:       sync.Mutex{},
	}
}

func (p *Proxy) Configuration() configuration.Configuration {
	return &p.config
}

func (p *Proxy) Setup(ctx context.Context) {
	p.client = kuaidailigo.NewAccountClient(
		p.config.SecretId,
		p.config.SecretKey,
		kuaidailigo.WithHttpClient(&http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		}),
	)
	p.debounce = debounce.New(time.Duration(p.config.DebounceTime))
}

func (p *Proxy) createProxyClient(ctx context.Context) error {
	p.proxyMutex.Lock()
	defer p.proxyMutex.Unlock()

	if p.proxyClient != nil {
		return nil
	}

	r, err := p.client.GetAccountOrders(ctx, kuaidailigo.GetAccountOrdersRequest{
		PayType: kuaidailigo.PayTypePostPay,
		Product: kuaidailigo.ProductTypeTPS,
		Status:  kuaidailigo.OrderStatusValid,
	})
	if err != nil || len(r) == 0 {
		p.Logger.Debug("create proxy client", log.Context(ctx))
		c, err := p.client.CreateOrder(ctx, kuaidailigo.CreateOrderRequest{
			IsNotify: false,
			PayType:  kuaidailigo.PayTypePostPay,
			CreateOrderParams: kuaidailigo.CreateTPSParams{
				Period:       0,
				MaxRps:       5,
				MaxBandwidth: 3,
				IPPool:       1,
			},
		})
		if err != nil {
			return err
		}
		p.proxyClient = &kuaidailigo.TPSOrderClient{OrderClient: c}
		time.Sleep(time.Second * 5) // 	// Wait for the proxy to be available
	} else {
		p.Logger.Debug("using existing proxy client", zap.String("order_id", r[0].OrderID), log.Context(ctx))
		c, err := p.client.GetOrderClient(ctx, r[0].OrderID)
		if err != nil {
			return err
		}
		p.proxyClient = &kuaidailigo.TPSOrderClient{OrderClient: c}
	}
	username, password, err := p.proxyClient.GetProxyAuthorization(ctx)
	if err != nil {
		return err
	}

	var u *url.URL

	for {
		u, err = p.proxyClient.GetProxy(ctx, kuaidailigo.ProxyProtocolHTTP)
		if err != nil {
			if strings.Contains(err.Error(), "407") { //一开始创建的时候会有一段时间查询不到
				p.Logger.Debug("getting proxy url failed, retrying", log.Context(ctx))
				time.Sleep(time.Second)
				continue
			}
			return err
		}
		break
	}

	u.User = url.UserPassword(username, password)
	p.proxyUrl = u
	p.Logger.Debug("create proxy client success", zap.String("proxy_url", p.proxyUrl.String()), log.Context(ctx))

	return nil
}

func (p *Proxy) closeProxyClient(ctx context.Context) {
	p.proxyMutex.Lock()
	defer p.proxyMutex.Unlock()

	if p.proxyClient == nil {
		return
	}
	p.debounce(func() {
		ctx, span := otel.GetTracerProvider().Tracer("github.com/yoshino-s/soar-helper/internal/proxy").Start(ctx, "close_proxy_client", trace.WithNewRoot())
		defer span.End()

		p.Logger.Debug("close proxy client")
		p.proxyClient.Close(ctx)
		p.Logger.Debug("close proxy client success")

		p.proxyClient = nil
		p.proxyUrl = nil
	})
}

func (p *Proxy) Do(ctx context.Context, f func(ctx context.Context, url *url.URL) error) error {
	if err := p.createProxyClient(ctx); err != nil {
		return err
	}
	defer p.closeProxyClient(ctx)

	p.Logger.Debug("do proxy request", zap.String("proxy_url", p.proxyUrl.String()), log.Context(ctx))
	return f(ctx, p.proxyUrl)
}
