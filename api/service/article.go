package service

import (
	"api/model"
	"api/repository"
)

type ArticleResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	LikeCount int    `json:"like_count"`
}

type ArticleService interface {
	Create(name, content string) (err error)
	List() (result []ArticleResponse, err error)
	Like(id int64) (err error)
}

func NewArticleService() ArticleService {
	return articleService{
		repository: repository.NewArticleRepository(),
	}
}

type articleService struct {
	repository repository.ArticleRepository
}

func (service articleService) Create(name, content string) (err error) {
	article := model.Article{
		Name:    name,
		Content: content,
	}
	service.repository.Insert(&article)
	if err != nil {
		return
	}
	return
}

func (service articleService) List() (result []ArticleResponse, err error) {

	// get articles
	articles, err := service.repository.List()
	if err != nil {
		return
	}

	// convert article model to article response
	for _, v := range articles {
		ar := ArticleResponse{
			ID:        v.ID,
			Name:      v.Name,
			Content:   v.Content,
			LikeCount: v.LikeCount,
		}
		result = append(result, ar)
	}

	return
}

func (service articleService) Like(id int64) (err error) {
	return
}
