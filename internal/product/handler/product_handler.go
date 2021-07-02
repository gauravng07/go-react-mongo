package handler

import (
	"encoding/json"
	"go-react-mongo/internal/logger"
	"go-react-mongo/internal/product/service"
	"net/http"
)

func ProductHandler(svc service.Listing) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := svc.GetListing(ctx)
		if err != nil {
			logger.Errorf(ctx, "error getting product details %+v", err)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
	return http.HandlerFunc(fn)
}

func BrandHandler(svc service.Listing) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := svc.GetListing(ctx)
		if err != nil {
			logger.Errorf(ctx, "error getting product details %+v", err)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
	return http.HandlerFunc(fn)
}
