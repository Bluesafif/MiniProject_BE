package Movie

import (
	"MiniProject_BE/dto/in"
	"MiniProject_BE/dto/out"
	"MiniProject_BE/model/applicationModel"
	"log"
	"net/http"
)

func (input movieService) GetListClientAssetOwner(request *http.Request) (response http.ResponseWriter, err error) {
	var (
		result        out.Response
		inputStruct   in.GetListDataDTO
		searchByParam []in.SearchByParam
	)

	inputStruct, searchByParam, err = input.ReadAndValidateGetListData(request, input.ValidSearchBy, input.ValidOrderBy, applicationModel.GetListMovieValidOperator, input.ValidLimit)
	if err.Error != nil {
		return
	}

	output, errs := input.doGetListMovie(inputStruct, searchByParam)
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output

	return
}

func (input movieService) doGetListMovie(inputStructInterface interface{}, timeNow []in.SearchByParam) (output interface{}, err error) {
	return
}
