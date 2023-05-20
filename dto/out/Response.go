package out

import "time"

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type DataInsert struct {
	Id         int64     `json:"id"`
	InsertedAt time.Time `json:"inserted_at"`
}

type DataUpdate struct {
	Id        int64     `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}
