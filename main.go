package main

import (
	// "fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.Handle("/assets/", http.FileServer(http.Dir(".")))
	serverStruct := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	serverStruct.ListenAndServe()
}
