package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/go-framework/telemetry"
	"github.com/yoshino-s/go-template/persistents/db"
)

var name = "go-template"
var app = application.NewMainApplication()

var (
	telemrtryApp = telemetry.New()
	entApp       = db.New()
	rootCmd      = &cobra.Command{
		Use: name,
	}
)

func init() {
	cobra.OnInitialize(func() {
		configuration.Setup(name)

		app.Append(telemrtryApp)
		app.Append(entApp)
	})

	configuration.GenerateConfiguration.Register(rootCmd.PersistentFlags())
	app.Configuration().Register(rootCmd.PersistentFlags())
	telemrtryApp.Configuration().Register(rootCmd.PersistentFlags())
	entApp.Configuration().Register(rootCmd.PersistentFlags())

	rootCmd.AddCommand(common.VersionCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
