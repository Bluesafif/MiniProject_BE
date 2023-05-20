package Movie

import (
	"MiniProject_BE/dto/in"
	"MiniProject_BE/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type movieService struct {
	service.GetListData
}

var MovieService = movieService{}.New()

func (input movieService) New() (output movieService) {
	return
}

func (input movieService) readBodyAndValidate(request *http.Request, validation func(input *in.Movie) error) (inputStruct in.Movie, errs error) {
	stringBody, _, err := service.ReadBody(request)
	if err != nil {
		return
	}

	_ = json.Unmarshal([]byte(stringBody), &inputStruct)

	id, _ := strconv.Atoi(mux.Vars(request)["ID"])
	if inputStruct.ID == 0 {
		inputStruct.ID = int64(id)
	}

	err = validation(&inputStruct)

	return
}
