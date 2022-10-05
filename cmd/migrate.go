package cmd

import (
	"fmt"
	"github.com/MelkoV/go-learn-admin/cmd/migrate"
	"github.com/MelkoV/go-learn-common/app"
	"github.com/MelkoV/go-learn-logger/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrate command
var (
	action string

	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			l := logger.NewCategoryLogger("admin/migrate", app.SYSTEM_UUID, logger.NewStreamLog())
			dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
				viper.GetString("db.user"),
				viper.GetString("db.password"),
				viper.GetString("db.host"),
				viper.GetString("db.port"),
				viper.GetString("db.name"),
			)
			switch action {
			case "up":
				l.Info("run migrations up")
				migrate.Up(l, viper.GetString("path.migrations"), dsn)
			default:
				l.Error("action not found")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")
	rootCmd.PersistentFlags().StringVar(&action, "action", "", "")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
