package chinaz

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
)

var _ configuration.Configuration = (*config)(nil)

type config struct {
	//Cookie string
	Token string
}

func (c *config) Register(set *pflag.FlagSet) {
	//set.String("chinaz.cookie", "", "chinaz cookie")
	set.String("chinaz.token", "", "chinaz token")
}

func (c *config) Read() {
	common.MustDecodeFromMapstructure(viper.AllSettings()["chinaz"], c)
}
