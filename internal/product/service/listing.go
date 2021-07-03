package service

import (
	"context"
	"go-react-mongo/internal/logger"
	"go-react-mongo/internal/product/model"
	"go-react-mongo/internal/product/repository"
)

type listingImpl struct {
	listingRepo repository.ListingRepo
}

func NewListingImpl(listingRepo repository.ListingRepo) *listingImpl {
	return &listingImpl{listingRepo: listingRepo}
}

type Listing interface {
	GetListing(ctx context.Context) ([]model.Listing, error)
	GetProductByBranch(ctx context.Context, brand string) ([]model.Listing, error)
	GetProductByPage(ctx context.Context, pageNum int, pageSize int) ([]model.Listing, error)
}

func (l *listingImpl) GetListing(ctx context.Context) ([]model.Listing, error){
	listing, err := l.listingRepo.GetAllProduct(ctx)
	if err != nil {
		logger.Errorf(ctx, "error getting products %+v", err)
	}
	return listing, err
}

func (l *listingImpl) GetProductByBranch(ctx context.Context, brand string) ([]model.Listing, error) {
	listing, err := l.listingRepo.GetProductByBrand(ctx, brand)
	if err != nil {
		logger.Errorf(ctx, "error getting product by brand %+v", err)
	}
	return listing, err
}

func (l *listingImpl) GetProductByPage(ctx context.Context, pageNum int, pageSize int) ([]model.Listing, error)  {
	listing, err := l.listingRepo.GetProductByPage(ctx, pageNum, pageSize)
	if err != nil {
		logger.Errorf(ctx, "error getting product by brand %+v", err)
	}
	return listing, err
}