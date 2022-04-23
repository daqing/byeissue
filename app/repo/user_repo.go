package repo

import (
	"github.com/daqing/byeissue/lib/orm"
	"gorm.io/gorm"
)

type UserRepo struct{}

func (u *UserRepo) CreateTable() {
	orm.DB().AutoMigrate(&userTable{})
}

func (u *UserRepo) Total() (count int64) {
	orm.DB().Model(&userTable{}).Count(&count)

	return
}

type userTable struct {
	gorm.Model
	Name string
}

func (userTable) TableName() string {
	return "users"
}
