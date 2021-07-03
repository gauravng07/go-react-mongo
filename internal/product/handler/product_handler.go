package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-react-mongo/internal/logger"
	"go-react-mongo/internal/product/service"
	"net/http"
	"strconv"
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

		vars := mux.Vars(r)
		brandName := vars["brand"]
		if len(brandName) == 0 {
			logger.Errorf(ctx, "error getting brand name")
			err := fmt.Errorf("error getting brand name")
			json.NewEncoder(w).Encode(err)
			return
		}

		response, err := svc.GetProductByBranch(ctx, brandName)
		if err != nil {
			logger.Errorf(ctx, "error getting product details %+v", err)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
	return http.HandlerFunc(fn)
}

func Search(svc service.Listing) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		queryParams := r.URL.Query()
		pageSize, err :=  strconv.Atoi(queryParams.Get("page_size"))
		if err != nil {
			logger.Errorf(ctx, "error invalid page size %+v", err)
			json.NewEncoder(w).Encode(err)
			return
		}
		pgNum, err := strconv.Atoi(queryParams.Get("page_num"))
		if err != nil {
			logger.Errorf(ctx, "error invalid page number %+v", err)
			json.NewEncoder(w).Encode(err)
			return
		}

		response, err := svc.GetProductByPage(ctx, pgNum, pageSize)
		if err != nil {
			logger.Errorf(ctx, "error getting product details %+v", err)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
	return http.HandlerFunc(fn)
}

