package ent

import (
	"github.com/yoshino-s/go-framework/extensions/ent"
	"go.uber.org/zap"
)

func (c *Client) DebugWithZap(l *zap.Logger) *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = ent.DebugEntDriver(c.driver, l)
	client := &Client{config: cfg}
	client.init()
	return client
}
