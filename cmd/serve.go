package cmd

import (
	"context"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/spf13/cobra"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/chinaz"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/handler/connect"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/handler/http"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/s3"
	"gitlab.yoshino-s.xyz/yoshino-s/soar-helper/utils"
)

func init() {
	httpApp.Configuration().Register(serveCmd.Flags())
	chinazApp.Configuration().Register(serveCmd.Flags())
	s3App.Configuration().Register(serveCmd.Flags())

	rootCmd.AddCommand(serveCmd)
}

var (
	chinazApp = chinaz.New()
	s3App     = s3.New()

	icpQueryService = connect.NewIcpQueryService()
	runService      = connect.NewRunnerService()
	toolsService    = connect.NewToolsService()
	s3Service       = connect.NewS3Service()

	httpApp = http.New()

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: `Serve runs the HTTP and gRPC server.`,
		Run: func(cmd *cobra.Command, args []string) {
			gologger.DefaultLogger.SetFormatter(utils.ToGoLoggerFormatter(app.Logger))
			gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose)

			app.Append(httpApp)
			app.Append(chinazApp)
			app.Append(s3App)
			app.Append(toolsService)

			chinazApp.SetDB(dbApp)

			icpQueryService.SetChinaz(chinazApp)
			icpQueryService.SetDB(dbApp)

			s3Service.SetS3(s3App)

			toolsService.SetS3(s3App)

			httpApp.SetIcpQueryHandler(icpQueryService)
			httpApp.SetRunnerHandler(runService)
			httpApp.SetToolsHandler(toolsService)
			httpApp.SetS3Handler(s3Service)

			app.Go(context.Background())
		},
	}
)
