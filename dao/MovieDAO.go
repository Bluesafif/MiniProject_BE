package dao

import (
	"MiniProject_BE/repository"
	"database/sql"
)

func InsertMovie(db *sql.Tx, userParam repository.MovieModel) (id int64, err error) {

	query := "INSERT INTO movie(title, description, duration, artists, genres, watch_url, created_at, updated_at) VALUES " +
		"($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
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
