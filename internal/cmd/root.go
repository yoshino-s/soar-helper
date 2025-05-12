package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-app/telemetry"
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
	dbApp        = db.New()
	telemetryApp = telemetry.New(
		context.Background(),
		telemetry.WithLogger(zap.NewExample()),
		telemetry.WithDSN("https://signoz-otl-http.yoshino-s.xyz"),
		telemetry.WithServiceName(name),
		telemetry.WithServiceVersion(common.Version),
	)
	proxyApp = proxy.New()
	rootCmd  = &cobra.Command{
		Use: name,
	}
)

func init() {
	common.AppName = name

	cobra.OnInitialize(func() {
		configuration.Setup(name)

		zap.ReplaceGlobals(app.Logger)
	})

	app.Append(dbApp)
	app.Append(proxyApp)
	app.Append(application.NewFuncApplication(application.StageClose, func(ctx context.Context) {
		telemetryApp.Close(ctx)
	}))

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
