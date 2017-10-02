package main

import (
	"github.com/gorilla/mux"
	"../go-restful-api-example/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/subjects").HandlerFunc(handler.GetCompanies)
	sub.Methods("POST").Path("/subjects").HandlerFunc(handler.SaveCompany)
	sub.Methods("GET").Path("/subjects/{name}").HandlerFunc(handler.GetCompany)
	sub.Methods("PUT").Path("/subjects/{name}").HandlerFunc(handler.UpdateCompany)
	sub.Methods("DELETE").Path("/subjects/{name}").HandlerFunc(handler.DeleteCompany)

	log.Fatal(http.ListenAndServe(":3000", router))
}
