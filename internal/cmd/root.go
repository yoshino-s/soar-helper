package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-app/otlp"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/cmd"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/soar-helper/internal/persistent/db"
	"github.com/yoshino-s/soar-helper/internal/proxy"
	"go.uber.org/zap"
)

var name = "soar-helper"
var app = application.NewMainApplication()

var (
	otlpApp = otlp.New(
		otlp.WithDSN("https://signoz-otl-http.yoshino-s.xyz/"),
		otlp.WithDeploymentEnvironment("Development"),
		otlp.WithServiceName(name),
		otlp.WithServiceVersion(common.Version),
	)
	dbApp    = db.New()
	proxyApp = proxy.New()
	rootCmd  = &cobra.Command{
		Use: name,
	}
)

func init() {
	common.AppName = name

	cobra.OnInitialize(func() {
		configuration.Setup(name)

		app.Append(otlpApp)
		app.Append(dbApp)
		app.Append(proxyApp)

		zap.ReplaceGlobals(app.Logger)
	})

	configuration.GenerateConfiguration.Register(rootCmd.PersistentFlags())
	app.Configuration().Register(rootCmd.PersistentFlags())
	dbApp.Configuration().Register(rootCmd.PersistentFlags())
	proxyApp.Configuration().Register(rootCmd.PersistentFlags())

	rootCmd.AddCommand(cmd.VersionCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
