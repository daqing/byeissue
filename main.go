package main

import (
	"net/http"

	"github.com/daqing/byeissue/app/repo"
	"github.com/daqing/byeissue/lib/orm"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/static", "./app/static")

	r.GET("/", func(c *gin.Context) {
		var count int64

		orm.DB().Model(&repo.UserTable{}).Count(&count)

		c.HTML(http.StatusOK, "home/index.html", gin.H{"Name": "it works", "Users": count})
	})

	r.Run(":2022")
}
