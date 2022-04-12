package main

import (
	get "./pkg/get"
	put "./pkg/put"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", get.GetStudents).Methods("GET")
	r.HandleFunc("/students/{name}", get.GetCities).Methods("GET")
	r.HandleFunc("/students/{name}/{city}", get.IsExists).Methods("GET")

	r.HandleFunc("/students/{name}", put.UpdateCities).Methods("PUT")
	r.HandleFunc("/students/{name}/{city}", put.UpdateExistance).Methods("PUT")
}
