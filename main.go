package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	srv := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
