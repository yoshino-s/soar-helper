package proxy

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/utils"
)

type config struct {
	SecretId     string        `mapstructure:"secret_id"`
	SecretKey    string        `mapstructure:"secret_key"`
	DebounceTime time.Duration `mapstructure:"debounce_time"`
}

func (c *config) Register(set *pflag.FlagSet) {
	set.String("proxy.secret_id", "", "kuaidaili account secret id")
	set.String("proxy.secret_key", "", "kuaidaili account secret key")
	set.Duration("proxy.debounce_time", 1*time.Minute, "order close debounce time")

	utils.MustNoError(viper.BindPFlags(set))
	configuration.Register(c)
}

func (c *config) Read() {
	utils.MustDecodeFromMapstructure(viper.AllSettings()["proxy"], c)
}
