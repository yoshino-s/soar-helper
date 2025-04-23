package beian

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/utils"
)

var _ configuration.Configuration = (*config)(nil)

type config struct {
	WerplusKey string `mapstructure:"werplus_key"`
	MiitSign   string `mapstructure:"miit_sign"`
}

func (c *config) Register(set *pflag.FlagSet) {
	//set.String("chinaz.cookie", "", "chinaz cookie")
	set.String("beian.werplus_key", "", "werplus key")
	set.String("beian.miit_sign", "", "sign token for miit")

	utils.MustNoError(viper.BindPFlags(set))
	configuration.Register(c)
}

func (c *config) Read() {
	utils.MustDecodeFromMapstructure(viper.AllSettings()["beian"], c)
}
