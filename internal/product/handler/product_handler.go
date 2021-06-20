package handler

import (
	"encoding/json"
	"net/http"
)

func ProductHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello")
	}
	return http.HandlerFunc(fn)
}
