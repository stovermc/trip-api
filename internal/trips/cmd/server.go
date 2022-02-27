package cmd

import (
	"context"
	"fmt"
	"os"

	kitlog "github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
	"github.com/stovermc/river-right-api/internal/trips/frameworks/persistence/mongo"
	"github.com/stovermc/river-right-api/internal/trips/frameworks/config"
	"github.com/stovermc/river-right-api/internal/trips/frameworks/log"
	"go.mongodb.org/mongo-driver/bson"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Run:   serverRun,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func serverRun(cmd *cobra.Command, args []string) {
	var kitLogger kitlog.Logger
	{
		kitLogger = kitlog.NewLogfmtLogger(os.Stderr)
		kitLogger = kitlog.With(kitLogger, "ts", kitlog.DefaultTimestampUTC)
		kitLogger = kitlog.With(kitLogger, "caller", kitlog.Caller(4))
	}

	logger := log.NewLogger(kitLogger)

	logger.Info("Yo yo we live!!")

	config := config.Init()

	conn := mongo.Connection{
		Username: config.MongoUsername,
		Password: config.MongoPassword,
		Host:     config.MongoHost,
		Port:     config.MongoPort,
		Datbase:  config.MongoDatabase,
	}

	mongoClient, disconnect, err := mongo.NewClient(logger, conn)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	defer disconnect()



	fmt.Println(mongoClient.ListDatabases(context.Background(), bson.M{}))
}
