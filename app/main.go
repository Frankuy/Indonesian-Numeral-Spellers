package main

import (
	"net/http"
	"log"
	routing "github.com/Indonesian-Numeral-Spellers/app/routing"
)

func main() {
	router := routing.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}