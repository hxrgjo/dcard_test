package repository

import (
	"api/model"
	"errors"

	"github.com/go-xorm/xorm"
)

type ArticleRepository interface {
	Insert(value *model.Article) (err error)
	List() (articles []model.Article, err error)
	Like(id int) (err error)
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{
		db: model.GetDB(),
	}
}

type articleRepository struct {
	db *xorm.Engine
}

func (a *articleRepository) Insert(value *model.Article) (err error) {
	_, err = a.db.InsertOne(value)
	if err != nil {
		return
	}
	return
}

func (a *articleRepository) List() (articles []model.Article, err error) {
	err = a.db.Desc("like_count").Asc("id").Find(&articles)
	if err != nil {
		return
	}
	return
}

func (a *articleRepository) Like(id int) (err error) {
	result, err := a.db.Exec("UPDATE articles SET like_count = like_count +1 WHERE id = ?", id)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected == 0 {
		return errors.New("data not found")
	}
	return
}
