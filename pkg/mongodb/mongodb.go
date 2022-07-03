package mdb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	_defaultConnectTimeout = time.Second
)

type Mongo struct {
	connectTimeout time.Duration

	Database *mongo.Database
}

func New(ctx context.Context, url string, database string, opts ...Option) (*Mongo, error) {
	// "mongodb://%s:%s@%s:%s"
	mdb := &Mongo{
		connectTimeout: _defaultConnectTimeout,
	}
	for _, opt := range opts {
		opt(mdb)
	}
	log.Printf("connect to Mongo by URL: %s", url)

	clientOptions := options.Client().ApplyURI(url)

	clientOptions.ConnectTimeout = &mdb.connectTimeout

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err)
	}
	mdb.Database = client.Database(database)
	return mdb, nil
}
