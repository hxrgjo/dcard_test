package main

import (
	"api/auth"
	"api/handler"
	"api/model"
	"api/repository"
	"api/service"

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

	// set jwt secret
	auth.SetSecret(config.JWTSecret)

	// run api server
	run()
}

func run() {
	// new user handler
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// new article handler
	articleRepository := repository.NewArticleRepository()
	articleService := service.NewArticleService(articleRepository)
	articleHandler := handler.NewArticleHandler(articleService)

	server := NewServer(userHandler, articleHandler)
	server.Run()
}
