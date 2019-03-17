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

type NewSignUpRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (handler *UserHandler) SignUp(c *gin.Context) {

	request := &NewSignUpRequest{}
	err := c.ShouldBindJSON(request)
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
			Message: "request fail",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
	})
}
