package product

import (
	"github.com/gorilla/mux"
	"go-react-mongo/internal/mongoDB"
	"go-react-mongo/internal/product/handler"
	"go-react-mongo/internal/product/repository"
	"go-react-mongo/internal/product/service"
	"net/http"
)

func Configure(productRouter *mux.Router, mongo *mongoDB.Client)  {
	listingRepo := repository.NewListingRepoImpl(mongo.MClient)
	listingSvc := service.NewListingImpl(listingRepo)
	r := productRouter.PathPrefix("/product").Subrouter()
	r.HandleFunc("/get-listing", handler.ProductHandler(listingSvc)).Methods(http.MethodGet)
}
