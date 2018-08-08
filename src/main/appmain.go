package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("\nServer Status: Initializing. \n")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8099", router))
}
