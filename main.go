package main

import (
	// "fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	serverStruct := http.Server{
		Handler: mux,
		Addr: ":8080",
	}
	serverStruct.ListenAndServe()
}