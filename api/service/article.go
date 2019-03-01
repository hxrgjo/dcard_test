package service

import (
	"api/model"
)

type articleResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	LikeCount int    `json:"like_count"`
}

type ArticleService interface {
	Create(name, content string) (err error)
	List() (result []articleResponse, err error)
	Like(id int64) (err error)
}

func NewArticleService() ArticleService {
	return articleService{}
}

type articleService struct {
}

func (service articleService) Create(name, content string) (err error) {
	article := model.Article{
		Name:    name,
		Content: content,
	}
	err = article.Insert()
	if err != nil {
		return
	}
	return
}

func (service articleService) List() (result []articleResponse, err error) {
	return
}

func (service articleService) Like(id int64) (err error) {
	return
}
