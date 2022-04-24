package home

import (
	"net/http"

	"github.com/daqing/byeissue/app/repo"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	r := &repo.UserRepo{}

	c.HTML(http.StatusOK, "home/index.html", gin.H{"Name": "it works", "Users": r.Total()})
}
