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
	productRouter.Use(applicationMiddleware)
	r := productRouter.PathPrefix("/product").Subrouter()
	r.HandleFunc("/get-listing", handler.ProductHandler(listingSvc)).Methods(http.MethodGet)
	r.HandleFunc("/search/brand/:brand", handler.BrandHandler(listingSvc)).Methods(http.MethodGet)
	r.HandleFunc("/search", handler.Search(listingSvc)).Methods(http.MethodGet)
}

func applicationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
