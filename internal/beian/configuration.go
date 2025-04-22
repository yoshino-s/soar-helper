package beian

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
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

	common.MustNoError(viper.BindPFlags(set))
	configuration.Register(c)
}

func (c *config) Read() {
	common.MustDecodeFromMapstructure(viper.AllSettings()["beian"], c)
}
