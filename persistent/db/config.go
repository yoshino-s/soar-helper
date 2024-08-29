package db

import (
	"entgo.io/ent/dialect"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
)

type Config struct {
	DriverName     string `mapstructure:"driver"`
	DataSourceName string `mapstructure:"datasource"`
}

var _ configuration.Configuration = (*Config)(nil)

func (c *Config) Register(flagSet *pflag.FlagSet) {
	flagSet.String("db.driver", dialect.SQLite, "database driver")
	flagSet.String("db.datasource", ":memory:?_pragma=foreign_keys(1)", "database datasource")
	common.MustNoError(viper.BindPFlags(flagSet))
	configuration.Register(c)
}

func (c *Config) Read() {
	common.MustDecodeFromMapstructure(viper.AllSettings()["db"], c)
}
