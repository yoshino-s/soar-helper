package db

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"entgo.io/ent/dialect/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-template/ent"
)

var _ application.Application = (*Client)(nil)

type Client struct {
	*ent.Client
	config Config
	logger *zap.Logger
}

var (
	client *Client
	mu     sync.RWMutex
)

func New() *Client {
	return &Client{
		config: Config{},
	}
}

func (c *Client) SetLogger(l *zap.Logger) {
	c.logger = l
}

func (c *Client) Configuration() configuration.Configuration {
	return &c.config
}

func (c *Client) Setup(context.Context) {
	drv := common.Must(sql.Open(c.config.DriverName, c.config.DataSourceName))
	c.Client = ent.NewClient(ent.Driver(drv))
	if c.config.Debug {
		c.Client = c.Client.DebugWithZap(c.logger)
	}
}
func (c *Client) Run(context.Context) {}
func (c *Client) Close(context.Context) {
	common.MustNoError(c.Client.Close())
}

func ReplaceGlobals(_client *Client) {
	mu.Lock()
	defer mu.Unlock()
	client = _client
}

func C() *Client {
	mu.RLock()
	defer mu.RUnlock()
	return client
}
