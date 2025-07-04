package db

import (
	"context"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/utils"
	"github.com/yoshino-s/soar-helper/internal/ent"
)

type Client struct {
	*application.EmptyApplication
	*ent.Client
	config Config
}

var (
	client *Client
	mu     sync.RWMutex
)

func New() *Client {
	return &Client{
		EmptyApplication: application.NewEmptyApplication("db"),
		config:           Config{},
	}
}

func (c *Client) Configuration() configuration.Configuration {
	return &c.config
}

func (c *Client) Initialize(context.Context) {
	drv := utils.Must(sql.Open(c.config.DriverName, c.config.DataSourceName))
	c.Client = ent.NewClient(ent.Driver(drv)).DebugWithZap(c.Logger)
}

func (c *Client) Close(context.Context) {
	utils.MustNoError(c.Client.Close())
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
