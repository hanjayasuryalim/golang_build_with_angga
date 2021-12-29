package main

import (
	"belajar/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/stores", handler.StoreHandler)
	mux.HandleFunc("/forms", handler.FormHandler)
	mux.HandleFunc("/process", handler.ProcessHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)

}
