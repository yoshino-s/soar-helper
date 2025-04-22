package beian

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
)

var _ configuration.Configuration = (*config)(nil)

type config struct {
	WerplusKey       string `mapstructure:"werplus_key"`
	IcpQueryEndpoint string `mapstructure:"icp_query_endpoint"`
}

func (c *config) Register(set *pflag.FlagSet) {
	//set.String("chinaz.cookie", "", "chinaz cookie")
	set.String("beian.werplus_key", "", "werplus key")
	set.String("beian.icp_query_endpoint", "", "icp query endpoint")

	common.MustNoError(viper.BindPFlags(set))
	configuration.Register(c)
}

func (c *config) Read() {
	common.MustDecodeFromMapstructure(viper.AllSettings()["beian"], c)
	c.IcpQueryEndpoint = viper.GetString("beian.icp_query_endpoint")
}
