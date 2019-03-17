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
		mock.EXPECT().Insert("test", "content").Return(nil)

		r := service.NewArticleServiceWithRepository(mock)
		err := r.Create("test", "content")
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

		r := service.NewArticleServiceWithRepository(mock)
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
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock.NewMockArticleRepository(ctrl)
		mock.EXPECT().Like(1).Return(nil)

		r := service.NewArticleServiceWithRepository(mock)
		err := r.Like(1)
		So(err, ShouldBeNil)
	})
}
