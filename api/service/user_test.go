package service_test

import (
	"api/auth"
	"api/model"
	"api/service"
	"testing"

	"api/mock"

	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSignUp(t *testing.T) {
	Convey("Test sign up", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock.NewMockUserRepository(ctrl)
		mock.EXPECT().CreateUser("test@gmail.com", gomock.Any(), "test")
		r := service.NewUserServiceWithRepository(mock)

		Convey("assert sign up result", func() {
			err := r.SignUp("test@gmail.com", "1234", "test")
			So(err, ShouldBeNil)
		})
	})
}

func TestSignIn(t *testing.T) {
	auth.SetSecret("test123")
	Convey("Test sign in", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock.NewMockUserRepository(ctrl)
		mock.EXPECT().FindByEmail("test@gmail.com").Return(
			model.User{
				ID:             1,
				Name:           "test",
				PasswordDigest: "$2a$10$lIWDJKBOtLSg6eRVwYosaeFwY6m5.1sH5vzovQctwPxBXynkxGSWS",
			},
			nil)
		r := service.NewUserServiceWithRepository(mock)

		Convey("assert sign in result", func() {
			_, err := r.SignIn("test@gmail.com", "1234")
			So(err, ShouldBeNil)
		})
	})
}
