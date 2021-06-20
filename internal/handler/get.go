package handler

import (
	"github.com/spf13/viper"
	"go-react-mongo/internal/config"
	"net/http"
)

func Get() http.Handler {
	return http.FileServer(http.Dir(viper.GetString(config.BuildDir)))
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}