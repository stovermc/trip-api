package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/stovermc/river-right-api/internal/trips/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Disconnect func() error

type Connection struct {
	Datbase  string
	Host     string
	Port     string
	Username string
	Password string
}

func (c *Connection) URI() string {
  return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", c.Username, c.Password, c.Host, c.Datbase)
}

func NewClient(logger app.Logger, conn Connection) (*mongo.Client, Disconnect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(conn.URI()))
	if err != nil {
		return nil, nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	logger.Info("testing mongo connection")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	disconnectFunc := func() error {
			logger.Info("disconnecting from mongo")
			return client.Disconnect(context.Background())
		}

	return client, disconnectFunc, nil
}
