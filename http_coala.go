package main

import (
	"log"
	"net/http"

	"http-coala/handlers"
)

func main() {
	http.HandleFunc("/coala/", handlers.GumHandler)
	http.HandleFunc("/cid/", handlers.CidHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
