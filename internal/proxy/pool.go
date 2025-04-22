package proxy

import (
	"context"
	"fmt"
	"net/url"
	"slices"
	"time"

	kuaidailigo "github.com/yoshino-s/kuaidaili-go"
	"go.uber.org/zap"
)

func (p *Proxy) Next(ctx context.Context) *url.URL {
	for {
		p.fill(ctx)
		index := p.current.Add(1)
		p := p.get(int(index % int32(len(p.proxies))))
		if p != nil {
			return p
		}
	}
}

func (p *Proxy) Invalidate(proxy *url.URL) {
	if proxy == nil {
		return
	}

	p.rwLock.Lock()
	defer p.rwLock.Unlock()

	p.proxies = slices.DeleteFunc(p.proxies, func(i *kuaidailigo.DPSProxy) bool {
		if i.Url == nil {
			fmt.Println("proxy url is nil???")
		}
		return i.Url.String() == proxy.String()
	})

	zap.L().Debug("proxies", zap.Any("proxies", p.proxies))
}

func (p *Proxy) get(idx int) *url.URL {
	p.rwLock.RLock()
	defer p.rwLock.RUnlock()

	if len(p.proxies) == 0 || idx < 0 || idx >= len(p.proxies) {
		return nil
	}

	i := p.proxies[idx]
	if i.Expire.Before(time.Now()) {
		zap.L().Debug("proxy expired", zap.Any("proxy", i))
		p.proxies = append(p.proxies[:idx], p.proxies[idx+1:]...)
		return nil
	}

	zap.L().Debug("proxies", zap.Any("proxies", p.proxies))

	return i.Url
}

func (p *Proxy) fill(ctx context.Context) error {
	p.rwLock.Lock()
	defer p.rwLock.Unlock()

	cnt := p.config.Count - int32(len(p.proxies))
	if cnt <= 0 {
		return nil
	}

	proxies, err := p.client.GetDPS(ctx, int(cnt), kuaidailigo.ProxyProtocolHTTP, nil)
	if err != nil {
		return err
	}
	for _, i := range proxies {
		if i.Url == nil {
			fmt.Println("proxy url is nil???")
		}
		zap.L().Debug("append proxy", zap.String("proxy", i.Url.String()), zap.Time("expire", i.Expire))
		p.proxies = append(p.proxies, i)
	}

	zap.L().Debug("proxies", zap.Any("proxies", p.proxies))

	return nil
}
