package dao

import (
	"MiniProject_BE/dto/in"
	"MiniProject_BE/repository"
	"database/sql"
)

func InsertMovie(db *sql.DB, userParam repository.MovieModel) (id int64, err error) {

	query := "INSERT INTO movie(title, description, duration, artists, genres, watch_url, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	param := []interface{}{
		userParam.Title.String, userParam.Description.String, userParam.Duration.String, userParam.Artists.String,
		userParam.Genres.String, userParam.WatchUrl.String, userParam.CreatedAt.Time, userParam.UpdatedAt.Time,
	}

	results := db.QueryRow(query, param...)
	errs := results.Scan(&id)
	if errs != nil && errs != sql.ErrNoRows {
		err = errs
		return
	}

	err = nil

	return
}

func UpdateMovie(db *sql.DB, userParam repository.MovieModel) (err error) {

	query :=
		" UPDATE movie " +
			" SET " +
			" title = $1, description = $2, duration = $3, " +
			" artists = $4, genres= $5, watch_url = $6, updated_at = $7 " +
			" WHERE " +
			" id = $8 "
	param := []interface{}{userParam.Title.String, userParam.Description.String, userParam.Duration.String, userParam.Artists.String,
		userParam.Genres.String, userParam.WatchUrl.String, userParam.UpdatedAt.Time, userParam.ID.Int64}

	stmt, errS := db.Prepare(query)
	defer stmt.Close()
	if errS != nil {
		err = errS
		return
	}

	_, errS = stmt.Exec(param...)
	if errS != nil {
		err = errS
		return
	}

	err = nil
	return
}

func GetListMovie(db *sql.DB, userParam in.GetListDataDTO, searchBy []in.SearchByParam) (result []interface{}, err error) {
	query :=
		"SELECT " +
			"	id, title, description, duration, artists, genres, watch_url, created_at, updated_at " +
			" FROM movie "

	return GetListDataDAO.GetListData(db, query, userParam, searchBy, 0, func(rows *sql.Rows) (interface{}, error) {
		var temp repository.MovieModel
		errors := rows.Scan(&temp.ID, &temp.Title, &temp.Description, &temp.Duration, &temp.Artists, &temp.Genres, &temp.WatchUrl, &temp.CreatedAt, &temp.UpdatedAt)
		return temp, errors
	}, "")
}
