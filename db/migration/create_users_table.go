package migration

import "github.com/daqing/byeissue/app/repo"

func CreateUsersTable() {
	r := &repo.UserRepo{}

	r.CreateTable()
}
