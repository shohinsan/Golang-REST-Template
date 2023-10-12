package myapi

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	message := "Welcome To Homepage!"
	_, err := w.Write([]byte(message))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data from the Api"
	_, err := w.Write([]byte(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
