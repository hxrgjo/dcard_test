package router

import (
	"api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	articleHandler := handler.NewArticleHandler()

	groupAPI := r.Group("/api")
	groupArticle := groupAPI.Group("/articles")
	groupArticle.POST("/", articleHandler.NewArticle)
	groupArticle.GET("/", articleHandler.GetArticles)
	groupArticle.PATCH("/:id/like", articleHandler.LikeArticle)
}
