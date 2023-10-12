package myapi

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	apiPrefix := "/api/data"

	apiMux := http.NewServeMux()
	mux.Handle(apiPrefix+"/", http.StripPrefix(apiPrefix, apiMux))

	apiMux.HandleFunc("/indexHandler", indexHandler)
	apiMux.HandleFunc("/apiDataHandler", apiDataHandler)

	return mux
}
