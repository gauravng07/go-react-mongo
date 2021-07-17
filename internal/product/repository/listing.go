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
	GetProductByPageOptimise(ctx context.Context, limit int, lastId string) ([]*model.Listing, error)
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
		return nil, fmt.Errorf("no product with given brand: %+v", err)
	}
	defer cursor.Close(ctx)

 	var listing []model.Listing
	if err = cursor.All(ctx, &listing); err != nil {
		return nil, fmt.Errorf("error getting product listing: %+v", err)
	}
	return listing, nil
}

func (l *listingRepoImpl) GetProductByPageOptimise(ctx context.Context, limit int, lastId string) ([]*model.Listing, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, internal.DefaultHttpTimeout)
	defer cancelFunc()

	pageSize := int64(limit)
	if len(lastId) == 0 {
		cursor, err := l.client.Database(viper.GetString(config.DBName)).
			Collection(viper.GetString(config.Collection)).
			Find(ctx, bson.D{}, &options.FindOptions{Limit: &pageSize, Sort: "_id"})

		if err != nil {
			return nil, fmt.Errorf("error paginating product page: %+v", err)
		}

		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				return
			}
		}(cursor, ctx)

		var listing []*model.Listing
		if err = cursor.All(ctx, &listing); err != nil {
			return nil, fmt.Errorf("error getting product listing: %+v", err)
		}
		return listing, nil

	} else {
		cursor, err := l.client.Database(viper.GetString(config.DBName)).
			Collection(viper.GetString(config.Collection)).
			Find(ctx, bson.M{"_id": bson.M{"$gt": lastId}}, &options.FindOptions{Limit: &pageSize})

		if err != nil {
			return nil, fmt.Errorf("error paginating product page: %+v", err)
		}

		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				return
			}
		}(cursor, ctx)

		var listing []*model.Listing
		if err = cursor.All(ctx, &listing); err != nil {
			return nil, fmt.Errorf("error getting product listing: %+v", err)
		}
		return listing, nil
	}
}