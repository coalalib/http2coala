package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"http-coala/handlers"
)

func main() {
	port := flag.Int("port", 8080, "Port number to listen on")
	flag.Parse()

	http.HandleFunc("/", handlers.Handler)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting server on port %d...\n", *port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
