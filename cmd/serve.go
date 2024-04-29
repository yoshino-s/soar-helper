package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-template/handler/grpc"
	"github.com/yoshino-s/go-template/handler/http"
)

func init() {
	httpApp.Configuration().Register(serveCmd.Flags())
	rootCmd.AddCommand(serveCmd)
}

var (
	httpApp  = http.New()
	grpcApp  = grpc.New()
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: `Serve runs the HTTP and gRPC server.`,
		Run: func(cmd *cobra.Command, args []string) {
			httpApp.SetGrpcServer(grpcApp)
			app.Append(httpApp)
			app.Go(context.Background())
		},
	}
)
