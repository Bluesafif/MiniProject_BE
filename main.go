package main

import (
	"MiniProject_BE/dao"
	"MiniProject_BE/model/applicationModel"
	"MiniProject_BE/router"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	dao.DBConnection()
	applicationModel.InitiateDefaultOperator()
	router.ApiController()
}
