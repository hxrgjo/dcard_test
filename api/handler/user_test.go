package handler

import (
	"api/mock"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSignUp(t *testing.T) {

	Convey("Test SignUp\n", t, func() {
		// mock user service
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock.NewMockUserService(ctrl)
		mock.EXPECT().SignUp("test@gmail.com", "1234", "test").Return(nil)

		// new user handler
		h := NewUserHandler(mock)

		// prepare router
		router := gin.Default()
		router.POST("/api/users", h.SignUp)

		// request
		reqeustBody := `
		{
			"name": "test",
			"password" : "1234",
			"email": "test@gmail.com"
		}
		`
		w := performRequest(router, "POST", "/api/users", []byte(reqeustBody))

		Convey("assert http status code", func() {
			So(w.Code, ShouldEqual, http.StatusOK)
		})

		Convey("assert response body", func() {
			var response map[string]interface{}
			err := json.Unmarshal([]byte(w.Body.String()), &response)
			So(err, ShouldBeNil)
			So(0, ShouldEqual, response["code"])
			So("success", ShouldEqual, response["message"])
		})
	})
}

func TestSignIn(t *testing.T) {

	Convey("Test SignIn\n", t, func() {
		// mock user service
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock.NewMockUserService(ctrl)
		mock.EXPECT().SignIn("test@gmail.com", "1234").Return("token", nil)

		// new user handler
		h := NewUserHandler(mock)

		// prepare router
		router := gin.Default()
		router.POST("/api/users/login", h.SignIn)

		// request
		reqeustBody := `
		{
			"password" : "1234",
			"email": "test@gmail.com"
		}
		`
		w := performRequest(router, "POST", "/api/users/login", []byte(reqeustBody))

		Convey("assert http status code", func() {
			So(w.Code, ShouldEqual, http.StatusOK)
		})

		Convey("assert response body", func() {
			var response map[string]interface{}
			err := json.Unmarshal([]byte(w.Body.String()), &response)
			So(err, ShouldBeNil)
			So(0, ShouldEqual, response["code"])
			So("success", ShouldEqual, response["message"])
		})
	})
}
