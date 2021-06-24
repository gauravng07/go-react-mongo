package mongoDB

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go-react-mongo/internal"
	"go-react-mongo/internal/config"
	"go-react-mongo/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	MClient	*mongo.Client
}

func NewMongoClient(ctx context.Context) (*Client, error) {
	client, err := createConnection(ctx)
	return &Client{MClient: client}, err
}

func createConnection(ctx context.Context) (*mongo.Client, error)  {
	ctx, cancel := context.WithTimeout(ctx, internal.DefaultMongoDBTimeout)
	defer cancel()

	URI := fmt.Sprintf("mongoDB+srv://%s:%s@%s/%s?w=majority",
		viper.GetString(config.UserName),
		viper.GetString(config.Password),
		viper.GetString(config.ClusterAddress),
		viper.GetString(config.DBName),
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		logger.Errorf(ctx, "error connecting to mongodb %v", err)
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Errorf(ctx, "error ping mongo %v", err)
		return nil, err
	}

	return client, nil
}
