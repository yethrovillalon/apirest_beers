package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	server := http.ListenAndServe(":8081", router)

	log.Fatal(server)
}
