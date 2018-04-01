package main

import (
	"log"
	"net/http"

	"../GoToRest/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/students").HandlerFunc(handler.GetStudents)
	sub.Methods("POST").Path("/students").HandlerFunc(handler.SaveStudent)
	sub.Methods("DELETE").Path("/student/{name}").HandlerFunc(handler.DeleteStudent)
	sub.Methods("GET").Path("/student/{name}").HandlerFunc(handler.GetStudent)
	// sub.Methods("GET").Path("/subjects/{name}").HandlerFunc(handler.GetCompany)
	//sub.Methods("PUT").Path("/subjects/{name}").HandlerFunc(handler.UpdateCompany)

	log.Fatal(http.ListenAndServe(":3000", router))
}
