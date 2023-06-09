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

func (input movieService) InsertMovie(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateInsert)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doInsertMovie(dao.DBConnection(), inputStruct, time.Now())
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output.(out.DataInsert)
	result.Message = "Data Berhasil Di Tambahkan"
	json.NewEncoder(response).Encode(result)
	return
}

func (input movieService) doInsertMovie(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Movie)
	var ids int64

	movieModel := repository.MovieModel{
		ID:          sql.NullInt64{Int64: inputStruct.ID},
		Title:       sql.NullString{String: inputStruct.Title},
		Description: sql.NullString{String: inputStruct.Description},
		Duration:    sql.NullString{String: inputStruct.Duration},
		Artists:     sql.NullString{String: inputStruct.Artists},
		Genres:      sql.NullString{String: inputStruct.Genres},
		WatchUrl:    sql.NullString{String: inputStruct.WatchUrl},
		CreatedAt:   sql.NullTime{Time: timeNow},
		UpdatedAt:   sql.NullTime{Time: timeNow},
	}

	ids, err = dao.InsertMovie(tx, movieModel)
	if err != nil {
		return
	}

	output = out.DataInsert{
		Id:         ids,
		InsertedAt: timeNow,
	}
	err = nil
	return
}

func (input movieService) validateInsert(inputStruct *in.Movie) error {
	return inputStruct.ValidateInsertMovie()
}
