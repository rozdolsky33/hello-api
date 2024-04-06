package main

import (
	"hello-api/handlers"
	"hello-api/handlers/rest"
	"hello-api/translation"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck) // <1>

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
