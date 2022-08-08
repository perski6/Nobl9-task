package main

import (
	"github.com/gorilla/mux"
	"github.com/ybbus/jsonrpc/v3"
	"log"
	"net/http"
	"perski6.com/nobl9-task/constants"
	"perski6.com/nobl9-task/handlers"
)

const RandomApiUrl = "https://api.random.org/json-rpc/4/invoke"

func main() {
	constants.LoadEnv()
	rpcClient := jsonrpc.NewClient(RandomApiUrl)

	r := mux.NewRouter()

	r.HandleFunc("/random/mean", func(w http.ResponseWriter, r *http.Request) {
		handlers.RandomsHandler(w, r, rpcClient)
	}).Methods(http.MethodGet).Queries("requests", "{requests:[0-9]+}", "length", "{length:[0-9]+}")

	log.Fatal(http.ListenAndServe(":8080", r))
}
