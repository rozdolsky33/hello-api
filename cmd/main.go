package main

import (
	"hello-api/handlers"
	"hello-api/handlers/rest"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct { // <6>
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
