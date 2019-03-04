package service

import (
	"api/model"
	"api/repository"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	Convey("Test create article\n", t, func() {
		// mock article repository
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := repository.NewMockArticleRepository(ctrl)
		article := model.Article{
			Name:    "test",
			Content: "content",
		}
		mock.EXPECT().Insert(&article).Return(nil)

		r := NewArticleServiceWithRepository(mock)
		err := r.Create("test", "content")
		So(err, ShouldBeNil)
	})
}

func TestList(t *testing.T) {
	Convey("Test list articles\n", t, func() {
		// mock article repository
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := repository.NewMockArticleRepository(ctrl)
		mockResult := []model.Article{{
			Name:    "test",
			Content: "content",
		}}
		mock.EXPECT().List().Return(mockResult, nil)

		r := NewArticleServiceWithRepository(mock)
		actual, err := r.List()
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, []ArticleResponse{{
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
		mock := repository.NewMockArticleRepository(ctrl)
		mock.EXPECT().Like(1).Return(nil)

		r := NewArticleServiceWithRepository(mock)
		err := r.Like(1)
		So(err, ShouldBeNil)
	})
}
