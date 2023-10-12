package main

import (
	"fmt"
	"golangProject/internal/routes/myapi"
	"net/http"
)

func main() {
	router := myapi.NewRouter()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
