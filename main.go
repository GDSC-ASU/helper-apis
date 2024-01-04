package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/quote", handleGetQuote)
	http.HandleFunc("/image", handleGetImage)
	http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir("./files"))))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
