package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/chinaz"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/handler/grpc"
	"gitlab.yoshino-s.xyz/yoshino-s/icp-lookup/handler/http"
)

func init() {
	httpApp.Configuration().Register(serveCmd.Flags())
	chinazApp.Configuration().Register(serveCmd.Flags())

	rootCmd.AddCommand(serveCmd)
}

var (
	httpApp   = http.New()
	grpcApp   = grpc.New()
	chinazApp = chinaz.New()
	serveCmd  = &cobra.Command{
		Use:   "serve",
		Short: `Serve runs the HTTP and gRPC server.`,
		Run: func(cmd *cobra.Command, args []string) {
			httpApp.SetGrpcServer(grpcApp)
			grpcApp.SetChinaz(chinazApp)
			grpcApp.SetDB(dbApp)

			chinazApp.SetDB(dbApp)
			app.Append(httpApp)
			app.Append(chinazApp)
			app.Go(context.Background())
		},
	}
)
