package endpoint

import (
	"MiniProject_BE/service/Movie"
	"net/http"
)

type consumerEndpoint struct {
	FileName string
}

var ConsumerEndpoint = consumerEndpoint{}.New()

func (input consumerEndpoint) New() (output consumerEndpoint) {
	output.FileName = "ConsumerEndpoint.go"
	return
}

func (input consumerEndpoint) ConsumerWithoutParam(response http.ResponseWriter, request *http.Request) {
	//funcName := "ConsumerWithoutParam"
	var err error
	if request.Method == "POST" {
		err = Movie.ConsumerService.InsertConsumer(request, response)
	} else if request.Method == "PUT" {
		//Movie.ConsumerService.UpdateConsumer
	} else if request.Method == "GET" {
		//Movie.ConsumerService.ViewConsumer
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
