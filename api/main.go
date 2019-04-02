package main

import (
	"api/auth"
	"api/handler"
	"api/model"
	"api/repository"
	"api/service"

	"go.uber.org/dig"

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

func run() error {
	container := dig.New()
	// new user handler
	container.Provide(repository.NewUserRepository, dig.As(new(repository.UserRepository)))
	container.Provide(service.NewUserService, dig.As(new(service.UserService)))
	container.Provide(handler.NewUserHandler)
	// new article handler
	container.Provide(repository.NewArticleRepository, dig.As(new(repository.ArticleRepository)))
	container.Provide(service.NewArticleService, dig.As(new(service.ArticleService)))
	container.Provide(handler.NewArticleHandler)

	err := container.Invoke(func(userHandler handler.UserHandler, articleHandler handler.ArticleHandler) {
		server := NewServer(userHandler, articleHandler)
		server.Run()
	})

	return err
}
