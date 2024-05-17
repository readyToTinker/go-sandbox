package main

import (
	"log"
	"net/http"

	"sandbox.com/main/webserver/handlers"
)

func main() {
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.New()))
}
