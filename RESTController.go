package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter() //Router

func startServer() {
	RouteHandlers()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func RouteHandlers() {
	r.HandleFunc("/api/students", getAllFromDB).Methods("GET")
	r.HandleFunc("/api/students/{id}", getWithId).Methods("GET")
	r.HandleFunc("/api/students/{fName},{email},{pNum}", insertIntoDB).Methods("POST")
	r.HandleFunc("/api/students/update/name/{id},{fName}", updateName).Methods("PUT")
	r.HandleFunc("/api/students/update/email/{id},{email}", updateEmail).Methods("PUT")
	r.HandleFunc("/api/students/update/phone/{id},{pNum}", updateNumber).Methods("PUT")
	r.HandleFunc("/api/students/{id}", deleteWithId).Methods("DELETE")
	r.HandleFunc("/api/students/", deleteAllFromDB).Methods("DELETE")
}
