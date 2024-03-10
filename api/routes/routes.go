package routes

import (
	"log"
	"net/http"
	"projectm/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/upload", handlers.CsvHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
