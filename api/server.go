package main

import (
	"api/handler"
	"api/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine         *gin.Engine
	userHandler    handler.UserHandler
	articleHandler handler.ArticleHandler
}

func NewServer(userHandler handler.UserHandler, articleHandler handler.ArticleHandler) *Server {
	return &Server{
		userHandler:    userHandler,
		articleHandler: articleHandler,
		engine:         gin.Default(),
	}
}

func (s *Server) SetHandler() {
	groupAPI := s.engine.Group("/api")
	groupUser := groupAPI.Group("/users")
	groupUser.POST("", s.userHandler.SignUp)
	groupUser.POST("/login", s.userHandler.SignIn)

	groupArticle := groupAPI.Group("/articles").Use(middleware.AuthMiddleware)
	groupArticle.POST("/", s.articleHandler.NewArticle)
	groupArticle.GET("/", s.articleHandler.GetArticles)
	groupArticle.PATCH("/:id/like", s.articleHandler.LikeArticle)
}

func (s *Server) Run() {
	s.SetHandler()
	s.engine.Run(fmt.Sprintf(":%d", config.Port)) // listen and serve on 0.0.0.0:8080
}
