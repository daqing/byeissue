package orm

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var once sync.Once

func DB() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=daqing dbname=byeissue sslmode=disable TimeZone=Asia/Shanghai"
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Failed to open postgres database with dsn: %s\n", dsn))
		}
	})

	return db
}
