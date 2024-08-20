package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/cmd"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/telemetry"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/persistent/db"
)

var name = "icp-lookup"
var app = application.NewMainApplication()

var (
	telemetryApp = telemetry.New()
	dbApp        = db.New()
	rootCmd      = &cobra.Command{
		Use: name,
	}
)

func init() {
	cobra.OnInitialize(func() {
		configuration.Setup(name)

		app.Append(telemetryApp)
		app.Append(dbApp)
	})

	configuration.GenerateConfiguration.Register(rootCmd.PersistentFlags())
	app.Configuration().Register(rootCmd.PersistentFlags())
	telemetryApp.Configuration().Register(rootCmd.PersistentFlags())
	dbApp.Configuration().Register(rootCmd.PersistentFlags())

	rootCmd.AddCommand(cmd.VersionCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
