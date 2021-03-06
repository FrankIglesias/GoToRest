package handler

// import (
// 	"encoding/json"
// 	"github.com/gorilla/mux"
// 	"../db"
// 	"../model"
// 	"io/ioutil"
// 	"net/http"
// )

// func GetCompany(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	name := vars["subject"]

// 	com, ok := db.FindBy(name)
// 	if !ok {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	bytes, err := json.Marshal(com)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	writeJsonResponse(w, bytes)
// }

// func UpdateCompany(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	name := vars["subect"]

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	com := new(model.Company)
// 	err = json.Unmarshal(body, com)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	db.Save(name, com)
// }

// func DeleteCompany(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	name := vars["subject"]

// 	db.Remove(name)
// 	w.WriteHeader(http.StatusNoContent)
// }

// func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Write(bytes)
// }
