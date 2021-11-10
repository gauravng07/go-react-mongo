package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go-react-mongo/internal"
	"go-react-mongo/internal/config"
	"go-react-mongo/internal/handler"
	"go-react-mongo/internal/logger"
	"go-react-mongo/internal/mongoDB"
	"go-react-mongo/internal/product"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultCorrelationId = "00000000.00000000"
)

var (
	ctx context.Context
)

func init()  {
	ctx = internal.SetContextWithValue(context.Background(), internal.ContextKeyCorrelationID, defaultCorrelationId)
}

func main() {

	client, err := mongoDB.NewMongoClient(ctx)
	if err != nil {
		logger.Fatalf(ctx, "error connecting to mongo client %v", err)
	}

	server := &http.Server{
		Addr: ":" + viper.GetString(config.Port),
		Handler: createRouter(client),
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
		logger.Infof(ctx, "server started at port: %v", config.Port)
	}()
	gracefulStop(server, client)
}

func createRouter(client *mongoDB.Client) *mux.Router {
	r := mux.NewRouter()
	r.Use(accessControlMiddleware)
	product.Configure(r, client)
	r.PathPrefix("/static").Handler(http.StripPrefix("/", handler.Get()))
	r.Handle("/metrics", promhttp.Handler())
	r.PathPrefix("/").HandlerFunc(handler.IndexHandler(viper.GetString(config.BuildDir) + "/index.html"))
	return r
}

func gracefulStop(server *http.Server, client *mongoDB.Client) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-stop
	if err := server.Shutdown(ctx); err != nil {
	} else {
		logger.Info(ctx, "server closed")
	}
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if req.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, req)
	})
}