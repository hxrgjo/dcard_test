package router

import (
	"api/auth"
	"api/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	articleHandler := handler.NewArticleHandler()
	userHandler := handler.NewUserHandler()

	groupAPI := r.Group("/api")
	groupArticle := groupAPI.Group("/articles")
	groupArticle.Use(authMiddleware)
	groupArticle.POST("/", articleHandler.NewArticle)
	groupArticle.GET("/", articleHandler.GetArticles)
	groupArticle.PATCH("/:id/like", articleHandler.LikeArticle)

	groupUser := groupAPI.Group("/users")
	groupUser.POST("", userHandler.SignUp)
	groupUser.POST("/login", userHandler.SignIn)
}

func authMiddleware(c *gin.Context) {
	userID, err := auth.Verify(strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	newAuthToken, err := auth.Sign(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Writer.Header().Set("Authorization", "Bearer "+newAuthToken)
	c.Set("user_id", userID)
	c.Next()
}
