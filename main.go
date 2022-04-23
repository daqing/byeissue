package main

import (
	"net/http"

	"github.com/daqing/byeissue/app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/static", "./app/static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", models.User{Name: "David Zhang"})
	})

	r.Run(":2022")
}
