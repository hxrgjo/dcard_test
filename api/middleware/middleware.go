package middleware

import (
	"api/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
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
