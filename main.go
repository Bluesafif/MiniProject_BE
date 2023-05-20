package main

import (
	"MiniProject_BE/model/applicationModel"
	"MiniProject_BE/router"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "paramadaksa"
	DB_NAME     = "movie"
)

func main() {
	DBConnection()
	applicationModel.InitiateDefaultOperator()
	router.ApiController()
}

func DbConn() {
	_, err := gorm.Open("mysql", "root:paramadaksa@/dataset?charset=utf8&parseTime=True")
	CheckError(err)

	if err == nil {
		log.Println("Connected!")
	}
}

func DBConnection() *sql.DB {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost port=5432", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	CheckError(err)
	if err == nil {
		log.Println("Connected!")
	}

	return db
}

func CheckError(err error) {
	if err != nil {
		log.Print("Connection Failed : ", err)
	}
}
