package config

import (
	"github.com/daqing/byeissue/app/controllers/home"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", home.Index)
}
