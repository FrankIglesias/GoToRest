package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"../db"
	"../model"
	"github.com/gorilla/mux"
	"github.com/thoas/go-funk"
)

func GetStudents(w http.ResponseWriter, _ *http.Request) {
	students := db.FindAll()

	bytes, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJSONResponse(w, bytes)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	student := new(model.Student)
	err = json.Unmarshal(body, student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, ok := db.FindBy(strconv.Itoa(student.ID))
	if ok {
		http.Error(w, "Student already exists", http.StatusForbidden)
	}

	db.Save(strconv.Itoa(student.ID), student)
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]
	db.Remove(studentID)
	w.WriteHeader(http.StatusNoContent)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	student, ok := db.FindBy(studentID)
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	bytes, err := json.Marshal(student)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	writeJSONResponse(w, bytes)
}

func GetStudentAverage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	result, ok := db.FindBy(studentID)
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	student := result.(*model.Student)
	averageMark := funk.Sum(funk.Map(student.Subjects, func(subject model.Subject) int {
		return subject.Mark
	})) / float64(len(student.Subjects))
	bytes, _ := json.Marshal(averageMark)
	writeJSONResponse(w, bytes)
}

func writeJSONResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
