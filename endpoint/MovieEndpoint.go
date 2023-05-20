package endpoint

import (
	"MiniProject_BE/service/Movie"
	"net/http"
)

type movieEndpoint struct {
	FileName string
}

var MovieEndpoint = movieEndpoint{}.New()

func (input movieEndpoint) New() (output movieEndpoint) {
	output.FileName = "MovieEndpoint.go"
	return
}

func (input movieEndpoint) MovieWithoutParam(response http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "POST" {
		err = Movie.MovieService.InsertMovie(request, response)
	} else if request.Method == "GET" {
		err = Movie.MovieService.GetListMovie(request, response)
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (input movieEndpoint) MovieWithParam(response http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "PUT" {
		err = Movie.MovieService.UpdateMovie(request, response)
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
