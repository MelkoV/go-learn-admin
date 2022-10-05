package cmd

import (
	"fmt"
	"github.com/MelkoV/go-learn-common/app"
	"github.com/MelkoV/go-learn-logger/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.NewCategoryLogger("admin/gen", app.SYSTEM_UUID, logger.NewStreamLog())
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
			viper.GetString("db.host"),
			viper.GetString("db.user"),
			viper.GetString("db.password"),
			viper.GetString("db.name"),
			viper.GetString("db.port"),
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		//sqlDb, _ := db.DB()
		//sqlDb.SetConnMaxLifetime(time.Hour)
		if err != nil {
			l.Fatal("error db connection: %v", err)
			os.Exit(1)
		}

		g := gen.NewGenerator(gen.Config{
			OutPath: "./gen",
			Mode:    gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		})

		g.UseDB(db)
		g.GenerateAllTable()
		g.Execute()
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
