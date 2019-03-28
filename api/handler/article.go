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

func NewArticleHandler(service service.ArticleService) ArticleHandler {
	return ArticleHandler{
		service: service,
	}
}

func (handler *ArticleHandler) NewArticle(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    errorcode.ValidateError,
			Message: "user not valid",
		})
		return
	}

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

	err = handler.service.Create(request.Name, request.Conetnet, userID.(int64))
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
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ValidateError,
			Message: "user not valid",
		})
		return
	}

	paramID := c.Param("id")
	articleID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ValidateError,
			Message: "validate error",
		})
		return
	}

	err = handler.service.Like(articleID, userID.(int64))
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
