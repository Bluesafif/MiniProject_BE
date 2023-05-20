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

func (input movieService) InsertConsumer(request *http.Request, writer http.ResponseWriter) (err error) {
	var (
		result out.Response
		//tx     *sql.Tx
	)

	_, errS := input.readBodyAndValidate(request, input.validateInsert)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	//output, errs := input.doInsertConsumer(tx, inputStruct, time.Now())
	//if errs != nil {
	//	log.Fatal(errs)
	//	return
	//}

	result.Code = 200
	result.Status = "OK"
	//result.Data = output
	json.NewEncoder(writer).Encode(result)
	return
}

func (input movieService) doInsertConsumer(tx *sql.Tx, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Movie)
	var id int64

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

	id, err = dao.InsertMovie(tx, movieModel)
	if err != nil {
		return
	}

	output = id
	err = nil
	return
}

func (input movieService) validateInsert(inputStruct *in.Movie) error {
	return inputStruct.ValidateInsertMovie()
}
