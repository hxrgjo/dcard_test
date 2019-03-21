package repository

import (
	"api/model"
	"errors"

	"github.com/go-xorm/xorm"
)

type ArticleRepository interface {
	Insert(name, content string, userID int64) (err error)
	List() (articles []model.Article, err error)
	Like(id int64, userID int64) (err error)
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{
		db: model.GetDB(),
	}
}

type articleRepository struct {
	db *xorm.Engine
}

func (a *articleRepository) Insert(name, content string, userID int64) (err error) {
	article := model.Article{Name: name, Content: content, UserID: userID}
	_, err = a.db.InsertOne(article)
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

func (a *articleRepository) Like(id int64, userID int64) (err error) {

	session := a.db.NewSession()
	defer session.Close()

	// begin transaction
	err = session.Begin()
	if err != nil {
		return
	}

	session.Rollback()

	result, err := session.Exec("UPDATE articles SET like_count = like_count +1 WHERE id = ?", id)
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

	affected, err = session.InsertOne(&model.ArticleLike{ArticleID: id, UserID: userID})
	if err != nil {
		return
	}
	if affected == 0 {
		return
	}

	// transaction commit
	session.Commit()

	return
}
