package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go-react-mongo/internal/config"
	"go-react-mongo/internal/handler"
	"go-react-mongo/internal/logger"
	"go-react-mongo/internal/product"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultCorrelationId = "00000000.00000000"
)

var ctx context.Context

func main() {

	server := &http.Server{
		Addr: ":" + viper.GetString(config.Port),
		Handler: createRouter(),
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	gracefulStop(server)
}

func createRouter() *mux.Router {
	r := mux.NewRouter()
	product.Configure(r)
	r.PathPrefix("/static").Handler(http.StripPrefix("/", handler.Get()))
	r.PathPrefix("/").HandlerFunc(handler.IndexHandler(viper.GetString(config.BuildDir) + "/index.html"))
	return r
}

func gracefulStop(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-stop
	if err := server.Shutdown(ctx); err != nil {
	} else {
		logger.Info(ctx, "server closed")
	}
}