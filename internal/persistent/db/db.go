package db

import (
	"context"
	"sync"

	"entgo.io/ent/dialect/sql"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/utils"
	"github.com/yoshino-s/soar-helper/internal/ent"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
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

func (c *Client) Setup(context.Context) {
	db := utils.Must(otelsql.Open(c.config.DriverName, c.config.DataSourceName, otelsql.WithAttributes(
		semconv.DBSystemSqlite,
	)))
	drv := sql.OpenDB(c.config.DriverName, db)
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
