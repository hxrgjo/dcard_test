package main

import (
	"api/model"
	"api/router"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var config Config

func main() {
	var err error

	// read config.yaml
	config, err = readConfig()
	if err != nil {
		panic("init config error:" + err.Error())
	}

	// init mysql db connect
	err = model.InitDB(
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.DBName,
	)
	if err != nil {
		panic("init db error:" + err.Error())
	}

	// init api server
	initServer()
}

func initServer() {
	r := gin.Default()
	router.Get(r)
	r.Run(fmt.Sprintf(":%d", config.Port)) // listen and serve on 0.0.0.0:8080
}
