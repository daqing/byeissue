package migration

import (
	"github.com/daqing/byeissue/app/repo"
	"github.com/daqing/byeissue/lib/orm"
)

func CreateUsersTable() {
	orm.DB().AutoMigrate(&repo.UserTable{})
}
