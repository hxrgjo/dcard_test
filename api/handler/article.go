package handler

import (
	"api/errorcode"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = 0
)

type ArticleHandler struct {
	Service service.ArticleService
}

func NewArticleHandler() ArticleHandler {
	return ArticleHandler{
		Service: service.NewArticleService(),
	}
}

type NewArticleRequest struct {
	Name      string `json:"name"`
	Conetnet  string `json:"content"`
	LikeCount int    `json:"like_count"`
}

func (handler *ArticleHandler) NewArticle(c *gin.Context) {

	request := &NewArticleRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ValidateError,
			Message: err.Error(),
		})
		return
	}

	err = handler.Service.Create(request.Name, request.Conetnet)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ArticleServiceError,
			Message: "request fail",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
	})
}

func (handler *ArticleHandler) GetArticles(c *gin.Context) {
	response, err := handler.Service.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ArticleServiceError,
			Message: "request fail",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Data:    response,
		Code:    CodeSuccess,
		Message: "success",
	})
}
