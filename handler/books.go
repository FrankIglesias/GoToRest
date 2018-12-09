package handler

import (
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/thoas/go-funk"
	"net/http"
	"io/ioutil"
	"os"
	"strconv"
	"../model"
)

func GetBooks(w http.ResponseWriter, _ *http.Request) {
	jsonFile, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	writeJSONResponse(w, byteValue)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonFile, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	bookId := vars["id"]
	var books []model.Book
	json.Unmarshal(byteValue, &books)

	book := funk.Find(books, func(x model.Book) bool {
	return strconv.Itoa(x.ID) == bookId
	})

	bytes, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	writeJSONResponse(w, bytes)
}