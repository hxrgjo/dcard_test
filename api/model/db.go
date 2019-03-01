package model

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func InitDB(user, password, host, dbName string) (err error) {
	connStr := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8",
		user, password, host, dbName,
	)

	engine, err = xorm.NewEngine("mysql", connStr)
	if err != nil {
		return
	}

	err = engine.DB().Ping()
	if err != nil {
		return
	}

	return
}
