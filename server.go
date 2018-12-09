package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"../go-to-rest/handler"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/books").HandlerFunc(handler.GetBooks)
	sub.Methods("GET").Path("/books/{id}").HandlerFunc(handler.GetBook)
	log.Fatal(http.ListenAndServe(":3000", router))
}
