package main

import (
	"github.com/daqing/byeissue/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/static", "./app/static")

	config.RegisterRoutes(r)

	r.Run(":2022")
}
