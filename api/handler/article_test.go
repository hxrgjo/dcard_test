package handler

import (
	"api/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewArticle(t *testing.T) {
	Convey("Test create article\n", t, func() {
		// mock article service
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := service.NewMockArticleService(ctrl)
		mock.EXPECT().Create("test", "jsonContentccc").Return(nil)

		// new article handler
		h := NewArticleHandlerWithService(mock)

		// prepare router
		router := gin.Default()
		router.POST("/api/articles", h.NewArticle)

		// request
		reqeustBody := `
		{
			"name": "test",
			"content" : "jsonContentccc"
		}
		`
		w := performRequest(router, "POST", "/api/articles", []byte(reqeustBody))

		// assert http status coe
		So(w.Code, ShouldEqual, http.StatusOK)

		// assert response body
		var response map[string]interface{}
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		So(err, ShouldBeNil)
		So(0, ShouldEqual, response["code"])
		So("success", ShouldEqual, response["message"])
	})
}
func TestGetArticles(t *testing.T) {
	Convey("Test get articles\n", t, func() {
		// mock article service
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := service.NewMockArticleService(ctrl)
		mockResponse := []service.ArticleResponse{{
			Name:      "11",
			Content:   "content",
			LikeCount: 0,
		}}
		mock.EXPECT().List().Return(mockResponse, nil)

		// new article handler
		h := NewArticleHandlerWithService(mock)

		// prepare router
		router := gin.Default()
		router.GET("/api/articles", h.GetArticles)

		Convey("perform request", func() {
			w := performRequest(router, "GET", "/api/articles", nil)

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

	})
}

func TestLikeArticle(t *testing.T) {
	Convey("Test like article\n", t, func() {
		// mock article service
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := service.NewMockArticleService(ctrl)
		mock.EXPECT().Like(1).Return(nil)

		// new article handler
		h := NewArticleHandlerWithService(mock)

		// prepare router
		router := gin.Default()
		router.PATCH("/api/articles/:id/like", h.LikeArticle)

		w := performRequest(router, "PATCH", "/api/articles/1/like", nil)

		// assert http status code
		So(w.Code, ShouldEqual, http.StatusOK)

		// assert response body
		var response map[string]interface{}
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		So(err, ShouldBeNil)
		So(0, ShouldEqual, response["code"])
		So("success", ShouldEqual, response["message"])
	})
}

func performRequest(r http.Handler, method, path string, requestBody []byte) *httptest.ResponseRecorder {
	b := bytes.NewBuffer(requestBody)
	req, _ := http.NewRequest(method, path, b)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
