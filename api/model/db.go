package model

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func InitDB(user, password, host, dbName string) (err error) {
	connStr := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8",
		user, password, host, dbName,
	)

	db, err = xorm.NewEngine("mysql", connStr)
	if err != nil {
		return
	}

	err = db.DB().Ping()
	if err != nil {
		return
	}

	return
}

func GetDB() (engine *xorm.Engine) {
	return db
}
