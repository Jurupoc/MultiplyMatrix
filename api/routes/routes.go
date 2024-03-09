package routes

import (
	"log"
	"net/http"
	"projectm/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.HelloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
