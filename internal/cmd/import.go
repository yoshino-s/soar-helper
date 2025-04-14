package cmd

import (
	"context"
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/common"
	"go.uber.org/zap"
)

func init() {
	importCmd.Flags().String("data", "data.csv", "data file path, .csv")

	rootCmd.AddCommand(importCmd)
}

var (
	importCmd = &cobra.Command{
		Use:   "import",
		Short: `Import data from file to database.`,
		Run: func(cmd *cobra.Command, args []string) {
			app.Append(application.FuncApplication(func(ctx context.Context) {
				csvFile := common.Must(cmd.Flags().GetString("data"))
				f := common.Must(os.Open(csvFile))
				defer f.Close()

				r := csv.NewReader(f)
				headers := common.Must(r.Read())
				app.Logger.Info("headers", zap.Any("headers", headers))

				for {
					record, err := r.Read()
					if err != nil {
						break
					}
					// do something
					t := common.Must(strconv.ParseInt(record[9], 10, 64))
					dbApp.Icp.Create().
						SetCity(record[0]).
						SetProvince(record[1]).
						SetCompany(record[2]).
						SetHost(record[4]).
						SetOwner(record[6]).
						SetPermit(record[7]).
						SetType(record[8]).
						SetCreatedAt(time.Unix(t/1000, 0)).
						SetUpdatedAt(time.Unix(t/1000, 0)).
						Save(ctx)
				}
			}))

			app.Go(context.Background())
		},
	}
)
