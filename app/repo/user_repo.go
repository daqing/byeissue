package repo

import "gorm.io/gorm"

type UserRepo struct{}

type UserTable struct {
	gorm.Model
	Name string
}

func (UserTable) TableName() string {
	return "users"
}
