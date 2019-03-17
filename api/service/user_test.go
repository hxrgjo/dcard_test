package service_test

import (
	"api/service"
	"testing"

	"api/mock"

	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {
	Convey("Test create user\n", t, func() {
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
