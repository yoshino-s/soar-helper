package cmd

import (
	"context"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/common"
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
			app.Append(application.FuncApplication(func(ctx context.Context) {
				opts := []schema.MigrateOption{
					schema.WithGlobalUniqueID(true),
				}

				if ok, _ := cmd.Flags().GetBool("dry-run"); ok {
					common.MustNoError(dbApp.Schema.WriteTo(ctx, cmd.OutOrStdout(), opts...))
				} else {
					common.MustNoError(dbApp.Schema.Create(ctx, opts...))
				}
			}))

			app.Go(context.Background())
		},
	}
)
