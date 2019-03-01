package main

import (
	"api/model"

	_ "github.com/go-sql-driver/mysql"
)

var config Config

func main() {
	var err error

	config, err = readConfig()
	if err != nil {
		panic("init config error:" + err.Error())
	}

	err = model.InitDB(
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.DBName,
	)
	if err != nil {
		panic("init db error:" + err.Error())
	}
}
