package repository

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go-react-mongo/internal/config"
	"go-react-mongo/internal/product/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type listingRepoImpl struct {
	client *mongo.Client
}

type ListingRepo interface {
	GetAllProduct(ctx context.Context) ([]model.Listing, error)
	GetProductByBrand(ctx context.Context, brand string)  ([]model.Listing, error)
}

func NewListingRepoImpl(client *mongo.Client) ListingRepo {
	return &listingRepoImpl{client: client}
}

func (l listingRepoImpl) GetAllProduct(ctx context.Context) ([]model.Listing, error) {
	cursor, err := l.client.Database(viper.GetString(config.DBName)).Collection(viper.GetString(config.Collection)).Find(ctx, nil, nil)
	defer cursor.Close(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting product listing: %+v", err)
	}

	var listing []model.Listing
	if err = cursor.All(ctx, &listing); err != nil {
		return nil, fmt.Errorf("error getting product listing: %+v", err)
	}

	return listing, nil
}

func (l listingRepoImpl) GetProductByBrand(ctx context.Context, brand string) ([]model.Listing, error) {
	panic("implement me")
}