package cmd

import (
	"context"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	migrateCmd.Flags().Bool("dry-run", false, "print the SQL statement without executing it")

	rootCmd.AddCommand(migrateCmd)
}

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: `Migrate runs the schema migration for the database.`,
		Run: func(cmd *cobra.Command, args []string) {
			app.Append(application.NewFuncApplication(application.StageRun, func(ctx context.Context) {
				opts := []schema.MigrateOption{}

				if ok, _ := cmd.Flags().GetBool("dry-run"); ok {
					utils.MustNoError(dbApp.Schema.WriteTo(ctx, cmd.OutOrStdout(), opts...))
				} else {
					utils.MustNoError(dbApp.Schema.Create(ctx, opts...))
				}
			}))

			ctx, span := otel.Tracer("github.com/yoshino-s/soar-helper").
				Start(context.Background(), "cmd.migrate", trace.WithNewRoot())
			defer span.End()

			app.Go(ctx)
		},
	}
)
