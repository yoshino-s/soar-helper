package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-framework/configuration"
	grpc_handler "github.com/yoshino-s/go-template/handler/grpc"
	http_handler "github.com/yoshino-s/go-template/handler/http"
)

func init() {
	configuration.HttpHandlerConfiguration.Register(serveCmd.Flags())

	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		gh, err := grpc_handler.NewGoTemplateServiceService()
		if err != nil {
			panic(err)
		}
		h, err := http_handler.NewHandler(
			configuration.HttpHandlerConfiguration.Config,
			gh,
		)
		if err != nil {
			panic(err)
		}

		app.Add(h)

		app.Go(context.Background())
	},
}
