package service

import (
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
	Like(id int) (err error)
}

func NewArticleService() ArticleService {
	return &articleService{
		repository: repository.NewArticleRepository(),
	}
}

func NewArticleServiceWithRepository(repository repository.ArticleRepository) ArticleService {
	return &articleService{
		repository: repository,
	}
}

type articleService struct {
	repository repository.ArticleRepository
}

func (service *articleService) Create(name, content string) (err error) {
	return service.repository.Insert(name, content)
}

func (service *articleService) List() (result []ArticleResponse, err error) {
	result = make([]ArticleResponse, 0)

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

func (service *articleService) Like(id int) (err error) {
	err = service.repository.Like(id)
	if err != nil {
		return
	}
	return
}
