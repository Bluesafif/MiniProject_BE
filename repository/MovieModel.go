package repository

import (
	"database/sql"
)

type MovieModel struct {
	ID          sql.NullInt64
	Title       sql.NullString
	Description sql.NullString
	Duration    sql.NullString
	Artists     sql.NullString
	Genres      sql.NullString
	WatchUrl    sql.NullString
	Deleted     sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
