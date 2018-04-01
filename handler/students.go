package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../db"
	"../model"
	"github.com/gorilla/mux"
)

func GetStudents(w http.ResponseWriter, _ *http.Request) {
	students := db.FindAll()

	bytes, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJSONResponse(w, bytes)
}

func SaveStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	student := new(model.Student)
	err = json.Unmarshal(body, student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Save(student.Name, student)
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	db.Remove(name)
	w.WriteHeader(http.StatusNoContent)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	student, ok := db.FindBy(name)

	if !ok {
		http.Error(w, "Not found", http.StatusInternalServerError)
	}

	bytes, err := json.Marshal(student)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	writeJSONResponse(w, bytes)
}

func writeJSONResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
