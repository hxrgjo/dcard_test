package service_test

import (
	"api/mock"
	"api/model"
	"api/service"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	Convey("Test create article\n", t, func() {
		// mock article repository
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock.NewMockArticleRepository(ctrl)
		mock.EXPECT().Insert("test", "content", int64(1)).Return(nil)

		r := service.NewArticleService(mock)
		err := r.Create("test", "content", 1)
		So(err, ShouldBeNil)
	})
}

func TestList(t *testing.T) {
	Convey("Test list articles\n", t, func() {
		// mock article repository
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock.NewMockArticleRepository(ctrl)
		mockResult := []model.Article{{
			Name:    "test",
			Content: "content",
		}}
		mock.EXPECT().List().Return(mockResult, nil)

		r := service.NewArticleService(mock)
		actual, err := r.List()
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, []service.ArticleResponse{{
			Name:    "test",
			Content: "content",
		}})
	})
}

func TestLike(t *testing.T) {
	Convey("Test like articles id 1\n", t, func() {
		// mock article repository
		// var userID int64
		// userID = 1
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock.NewMockArticleRepository(ctrl)
		mock.EXPECT().Like(int64(1), int64(1)).Return(nil)

		r := service.NewArticleService(mock)
		err := r.Like(1, 1)
		So(err, ShouldBeNil)
	})
}
