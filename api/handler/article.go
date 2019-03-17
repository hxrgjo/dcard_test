package handler

import (
	"api/errorcode"
	"api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	service service.ArticleService
}

func NewArticleHandler() ArticleHandler {
	return ArticleHandler{
		service: service.NewArticleService(),
	}
}

func NewArticleHandlerWithService(service service.ArticleService) ArticleHandler {
	return ArticleHandler{
		service: service,
	}
}

func (handler *ArticleHandler) NewArticle(c *gin.Context) {
	request := struct {
		Name      string `json:"name"`
		Conetnet  string `json:"content"`
		LikeCount int    `json:"like_count"`
	}{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ValidateError,
			Message: err.Error(),
		})
		return
	}

	err = handler.service.Create(request.Name, request.Conetnet)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ArticleServiceError,
			Message: "create failed",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
	})
}

func (handler *ArticleHandler) GetArticles(c *gin.Context) {
	response, err := handler.service.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ArticleServiceError,
			Message: "get list failed",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Data:    response,
		Code:    CodeSuccess,
		Message: "success",
	})
}

func (handler *ArticleHandler) LikeArticle(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ValidateError,
			Message: "validate error",
		})
		return
	}

	err = handler.service.Like(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ArticleServiceError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
	})
}
