package Movie

import (
	"MiniProject_BE/dao"
	"MiniProject_BE/dto/in"
	"MiniProject_BE/dto/out"
	"MiniProject_BE/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (input movieService) UpdateMovie(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateInsert)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doUpdateMovie(dao.DBConnection(), inputStruct, time.Now())
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output.(out.DataUpdate)
	result.Message = "Data Berhasil Di Ubah"
	json.NewEncoder(response).Encode(result)
	return
}

func (input movieService) doUpdateMovie(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Movie)

	movieModel := repository.MovieModel{
		ID:          sql.NullInt64{Int64: inputStruct.ID},
		Title:       sql.NullString{String: inputStruct.Title},
		Description: sql.NullString{String: inputStruct.Description},
		Duration:    sql.NullString{String: inputStruct.Duration},
		Artists:     sql.NullString{String: inputStruct.Artists},
		Genres:      sql.NullString{String: inputStruct.Genres},
		WatchUrl:    sql.NullString{String: inputStruct.WatchUrl},
		UpdatedAt:   sql.NullTime{Time: timeNow},
	}

	err = dao.UpdateMovie(tx, movieModel)
	if err != nil {
		return
	}

	output = out.DataUpdate{
		Id:        inputStruct.ID,
		UpdatedAt: timeNow,
	}
	err = nil
	return
}

func (input movieService) validateUpdate(inputStruct *in.Movie) error {
	return inputStruct.ValidateUpdateMovie()
}
