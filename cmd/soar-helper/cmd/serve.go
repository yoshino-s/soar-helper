package cmd

import (
	"context"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/spf13/cobra"
	"github.com/yoshino-s/soar-helper/internal/handler/connect"
	"github.com/yoshino-s/soar-helper/internal/handler/http"
	"github.com/yoshino-s/soar-helper/internal/s3"
	"github.com/yoshino-s/soar-helper/internal/utils"
)

func init() {
	httpApp.Configuration().Register(serveCmd.Flags())
	s3App.Configuration().Register(serveCmd.Flags())

	rootCmd.AddCommand(serveCmd)
}

var (
	s3App = s3.New()

	httpApp = http.New()

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: `Serve runs the HTTP and gRPC server.`,
		Run: func(cmd *cobra.Command, args []string) {
			gologger.DefaultLogger.SetFormatter(utils.ToGoLoggerFormatter(app.Logger))
			gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose)

			app.Append(httpApp)
			app.Append(s3App)
			app.Append(connect.NewS3ServiceHandler())
			app.Append(connect.NewToolsServiceHandler())
			app.Append(connect.NewRunnerServiceHandler())

			app.Go(context.Background())
		},
	}
)
