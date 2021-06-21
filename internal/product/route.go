package product

import (
	"github.com/gorilla/mux"
	"go-react-mongo/internal/mongoDB"
	"go-react-mongo/internal/product/handler"
	"net/http"
)

func Configure(productRouter *mux.Router, mongoCli *mongoDB.Client)  {
	r := productRouter.PathPrefix("/product").Subrouter()
	r.HandleFunc("/get", handler.ProductHandler("test")).Methods(http.MethodGet)
}
