package out

import "time"

type MovieResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    string    `json:"duration"`
	Artists     string    `json:"artists"`
	Genres      string    `json:"genres"`
	WatchUrl    string    `json:"watch_url"`
	Deleted     string    `json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
