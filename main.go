package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/quote", handleGetQuote)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
