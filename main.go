package main

import (
	"net/http"

	"github.com/daqing/byeissue/app/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/static", "./app/static")

	r.GET("/", func(c *gin.Context) {
		r := &repo.UserRepo{}

		c.HTML(http.StatusOK, "home/index.html", gin.H{"Name": "it works", "Users": r.Total()})
	})

	r.Run(":2022")
}
