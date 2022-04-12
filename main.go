package main

import (
	"github.com/gorilla/mux"
	"github.com/vovovoC/hello/pkg/get"
	"github.com/vovovoC/hello/pkg/put"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", get.GetStudents).Methods("GET")
	r.HandleFunc("/students/{name}", get.GetCities).Methods("GET")
	r.HandleFunc("/students/{name}/{city}", get.IsExists).Methods("GET")

	r.HandleFunc("/students/{name}", put.UpdateCities).Methods("PUT")
	r.HandleFunc("/students/{name}/{city}", put.UpdateExistance).Methods("PUT")
}
