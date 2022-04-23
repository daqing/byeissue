package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/static", "./app/static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":   "Homepage",
			"message": "Hello, world from gin.H",
		})
	})

	r.Run(":2022")
}
