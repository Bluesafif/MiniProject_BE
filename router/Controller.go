package router

import (
	"MiniProject_BE/endpoint"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ApiController() {
	handle := mux.NewRouter()

	//insert
	handle.HandleFunc("/consumer/insert", endpoint.ConsumerEndpoint.ConsumerWithoutParam).Methods("POST")

	log.Fatal(http.ListenAndServe(":9999", handle))
}
