package beian

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
)

var _ configuration.Configuration = (*config)(nil)

type config struct {
	//Cookie string
	ChinazToken string `mapstructure:"chinaz_token"`
	WerplusKey  string `mapstructure:"werplus_key"`
}

func (c *config) Register(set *pflag.FlagSet) {
	//set.String("chinaz.cookie", "", "chinaz cookie")
	set.String("beian.chinaz_token", "", "chinaz token")
	set.String("beian.werplus_key", "", "werplus key")

	common.MustNoError(viper.BindPFlags(set))
	configuration.Register(c)
}

func (c *config) Read() {
	common.MustDecodeFromMapstructure(viper.AllSettings()["beian"], c)
}
