package model

import "time"

type ArticleLike struct {
	ID        int64      `xorm:"'id' pk autoincr"`
	UserID    int64      `xorm:"'user_id'"`
	ArticleID int64      `xorm:"'article_id'"`
	CreatedAt *time.Time `xorm:"created"`
}

func (ArticleLike) TableName() string {
	return "article_likes"
}
