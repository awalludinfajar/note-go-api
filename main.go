package main

import (
	"log"
	"net/http"

	"github.com/awalludinfajar/note-go-api.git/route"
)

func main() {
	router := route.SetupRoutes()

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
