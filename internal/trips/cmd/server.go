package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	kitlog "github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
	"github.com/stovermc/river-right-api/internal/trips/frameworks/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	uri := "mongodb+srv://stovermc:xH8QgPnSRu0FOs5A@cluster0.lzf8u.mongodb.net/river-right?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Error(err, "error creating mongo client")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		logger.Error(err, "error connecting to mongo")
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Error(err, "mongo ping failed")
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(databases)
}
