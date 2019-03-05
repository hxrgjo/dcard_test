package model

import (
	"fmt"
	"time"

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

	for {
		err := db.DB().Ping()
		if err != nil {
			fmt.Println("sleep 1 second for db init")
			time.Sleep(1 * time.Second)
			continue
		}
		return nil
	}
}

func GetDB() (engine *xorm.Engine) {
	return db
}
