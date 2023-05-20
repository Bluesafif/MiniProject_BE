package Movie

import (
	"MiniProject_BE/dao"
	"MiniProject_BE/dto/in"
	"MiniProject_BE/dto/out"
	"MiniProject_BE/model/applicationModel"
	"MiniProject_BE/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func (input movieService) GetListMovie(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result        out.Response
		inputStruct   in.GetListDataDTO
		searchByParam []in.SearchByParam
	)

	inputStruct, searchByParam, err = input.ReadAndValidateGetListData(request, input.ValidSearchBy, input.ValidOrderBy, applicationModel.GetListMovieValidOperator, input.ValidLimit)
	if err != nil {
		return
	}

	output, errs := input.doGetListMovie(dao.DBConnection(), inputStruct, searchByParam)
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output
	result.Message = "Pengambilan Data Berhasil"
	json.NewEncoder(response).Encode(result)

	return
}

func (input movieService) doGetListMovie(db *sql.DB, inputStruct in.GetListDataDTO, searchByParam []in.SearchByParam) (output []out.MovieResponse, err error) {
	var dbResult []interface{}

	dbResult, err = dao.GetListMovie(db, inputStruct, searchByParam)
	if err != nil {
		return
	}

	output = input.convertToListDTOOut(dbResult)
	return
}

func (input movieService) convertToListDTOOut(dbResult []interface{}) (result []out.MovieResponse) {
	for i := 0; i < len(dbResult); i++ {
		repo := dbResult[i].(repository.MovieModel)
		result = append(result, out.MovieResponse{
			ID:          repo.ID.Int64,
			Title:       repo.Title.String,
			Description: repo.Description.String,
			Duration:    repo.Duration.String,
			Artists:     repo.Artists.String,
			Genres:      repo.Genres.String,
			WatchUrl:    repo.WatchUrl.String,
			Deleted:     repo.Deleted.String,
			CreatedAt:   repo.CreatedAt.Time,
			UpdatedAt:   repo.UpdatedAt.Time,
		})
	}
	return result
}
