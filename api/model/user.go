package model

import "time"

type User struct {
	ID             int64 `xorm:"'id' pk autoincr"`
	Email          string
	PasswordDigest string
	Name           string
	CreatedAt      time.Time `xorm:"created"`
	UpdatedAt      time.Time `xorm:"updated"`
}

func (User) TableName() string {
	return "users"
}
