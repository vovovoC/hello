package main

import (
	"pkg/get"
	"pkg/put"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", get.GetStudents).Methods("GET")
	r.HandleFunc("/students/{name}", get.GetCities).Methods("GET")

	r.HandleFunc("/students/{name}", put.Save).Methods("PUT")
}
