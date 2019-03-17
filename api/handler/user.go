package handler

import (
	"api/errorcode"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler() UserHandler {
	return UserHandler{
		service: service.NewUserService(),
	}
}

func NewUserHandlerWithService(service service.UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (handler *UserHandler) SignUp(c *gin.Context) {
	request := struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ValidateError,
			Message: err.Error(),
		})
		return
	}

	err = handler.service.SignUp(request.Email, request.Password, request.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ArticleServiceError,
			Message: "sign up failed",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
	})
}

func (handler *UserHandler) SignIn(c *gin.Context) {
	request := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ValidateError,
			Message: err.Error(),
		})
		return
	}

	token, err := handler.service.SignIn(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errorcode.ArticleServiceError,
			Message: "sign in failed",
		})
		return
	}

	c.Writer.Header().Set("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
	})
}
