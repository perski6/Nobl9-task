package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"perski6.com/nobl9-task/constants"
	"perski6.com/nobl9-task/handlers"
)

func main() {
	constants.LoadEnv()
	r := mux.NewRouter()
	r.HandleFunc("/random/mean", handlers.RandomsHandler).Methods(http.MethodGet).Queries("requests", "{requests:[0-9]+}", "length", "{length:[0-9]+}")
	log.Fatal(http.ListenAndServe(":8080", r))
}
