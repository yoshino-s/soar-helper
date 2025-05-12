package cmd

import (
	"context"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/spf13/cobra"
	"github.com/yoshino-s/soar-helper/internal/beian"
	"github.com/yoshino-s/soar-helper/internal/handler/connect"
	"github.com/yoshino-s/soar-helper/internal/handler/http"
	"github.com/yoshino-s/soar-helper/internal/s3"
	"github.com/yoshino-s/soar-helper/internal/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	httpApp.Configuration().Register(serveCmd.Flags())
	beianApp.Configuration().Register(serveCmd.Flags())
	s3App.Configuration().Register(serveCmd.Flags())

	rootCmd.AddCommand(serveCmd)
}

var (
	beianApp = beian.New()
	s3App    = s3.New()

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
			app.Append(beianApp)
			app.Append(s3App)
			app.Append(toolsService)

			beianApp.Set(dbApp, proxyApp)

			icpQueryService.SetChinaz(beianApp)
			icpQueryService.SetDB(dbApp)

			s3Service.SetS3(s3App)

			toolsService.SetS3(s3App)

			httpApp.SetIcpQueryHandler(icpQueryService)
			httpApp.SetRunnerHandler(runService)
			httpApp.SetToolsHandler(toolsService)
			httpApp.SetS3Handler(s3Service)

			ctx, span := otel.Tracer("github.com/yoshino-s/soar-helper").
				Start(context.Background(), "cmd.serve", trace.WithNewRoot())
			defer span.End()

			app.Go(ctx)
		},
	}
)
