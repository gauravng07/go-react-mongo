package repository

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go-react-mongo/internal"
	"go-react-mongo/internal/config"
	"go-react-mongo/internal/product/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type listingRepoImpl struct {
	client *mongo.Client
}

type ListingRepo interface {
	GetAllProduct(ctx context.Context) ([]model.Listing, error)
	GetProductByBrand(ctx context.Context, brand string)  ([]model.Listing, error)
	GetProductByPage(ctx context.Context, pageNum int, pageSize int) ([]model.Listing, error)
}

func NewListingRepoImpl(client *mongo.Client) ListingRepo {
	return &listingRepoImpl{client: client}
}

const DefaultPageSize = 50

func (l *listingRepoImpl) GetAllProduct(ctx context.Context) ([]model.Listing, error) {

	ctx, cancelFunc := context.WithTimeout(ctx, internal.DefaultHttpTimeout)
	defer cancelFunc()

	cursor, err := l.client.Database(viper.GetString(config.DBName)).
		Collection(viper.GetString(config.Collection)).
		Find(ctx, bson.D{{}})

	if err != nil {
		return nil, fmt.Errorf("error getting product listing: %+v", err)
	}
	defer cursor.Close(ctx)
	var listing []model.Listing
	if err = cursor.All(ctx, &listing); err != nil {
		return nil, fmt.Errorf("error getting product listing: %+v", err)
	}

	return listing, nil
}

func (l *listingRepoImpl) GetProductByPage(ctx context.Context, pageNum int, pageSize int) ([]model.Listing, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, internal.DefaultHttpTimeout)
	defer cancelFunc()

	if pageSize == 0 {
		pageSize = DefaultPageSize
	}

	skip := int64(pageSize * (pageNum - 1))
	limit := int64(pageSize)

	cursor, err := l.client.Database(viper.GetString(config.DBName)).
		Collection(viper.GetString(config.Collection)).
		Find(ctx, bson.D{{}}, &options.FindOptions{Skip: &skip, Limit: &limit})

	if err != nil {
		return nil,  fmt.Errorf("error getting product by page: %+v", err)
	}
	defer cursor.Close(ctx)
	var listing []model.Listing
	if err = cursor.All(ctx, &listing); err != nil {
		return nil, fmt.Errorf("error getting product listing: %+v", err)
	}

	return listing, nil
}

func (l *listingRepoImpl) GetProductByBrand(ctx context.Context, brand string) ([]model.Listing, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, internal.DefaultHttpTimeout)
	defer cancelFunc()

	cursor, err := l.client.Database(viper.GetString(config.DBName)).
		Collection(viper.GetString(config.Collection)).
		Find(ctx, map[string]string{
			"brand": brand,
		}, nil)
	if err != nil {
		return nil, fmt.Errorf("no product woth given brancd: %+v", err)
	}
	defer cursor.Close(ctx)

 	var listing []model.Listing
	if err = cursor.All(ctx, &listing); err != nil {
		return nil, fmt.Errorf("error getting product listing: %+v", err)
	}
	return listing, nil
}