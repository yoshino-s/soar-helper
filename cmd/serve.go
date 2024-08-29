package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/chinaz"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/handler/connect"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/handler/http"
)

func init() {
	httpApp.Configuration().Register(serveCmd.Flags())
	chinazApp.Configuration().Register(serveCmd.Flags())

	rootCmd.AddCommand(serveCmd)
}

var (
	httpApp         = http.New()
	icpQueryService = connect.NewIcpQueryService()
	runService      = connect.NewRunnerService()
	chinazApp       = chinaz.New()
	serveCmd        = &cobra.Command{
		Use:   "serve",
		Short: `Serve runs the HTTP and gRPC server.`,
		Run: func(cmd *cobra.Command, args []string) {
			chinazApp.SetDB(dbApp)
			icpQueryService.SetChinaz(chinazApp)
			icpQueryService.SetDB(dbApp)

			httpApp.SetIcpQueryHandler(icpQueryService)
			httpApp.SetRunnerHandler(runService)

			app.Append(httpApp)
			app.Append(chinazApp)
			app.Go(context.Background())
		},
	}
)
