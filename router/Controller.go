package router

import (
	"MiniProject_BE/endpoint"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ApiController() {
	handle := mux.NewRouter()

	handle.HandleFunc("/movie", endpoint.MovieEndpoint.MovieWithoutParam).Methods("POST", "GET")
	handle.HandleFunc("/movie/{ID}", endpoint.MovieEndpoint.MovieWithParam).Methods("PUT")

	log.Fatal(http.ListenAndServe(":9999", handle))
}
