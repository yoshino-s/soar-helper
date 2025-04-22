package proxy

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
)

type config struct {
	SecretId  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`
	Count     int32  `mapstructure:"count"`
}

func (c *config) Register(set *pflag.FlagSet) {
	set.String("proxy.secret_id", "", "proxy secret id")
	set.String("proxy.secret_key", "", "proxy secret key")
	set.Int32("proxy.count", 0, "proxy count")

	common.MustNoError(viper.BindPFlags(set))
	configuration.Register(c)
}

func (c *config) Read() {
	common.MustDecodeFromMapstructure(viper.AllSettings()["proxy"], c)
}
