package controller

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", indexCtl)
	r.GET("/article/:id", articleCtl)
	return r
}

func indexCtl(c *gin.Context) {
	c.JSON(200,"hello")
}

func articleCtl(c *gin.Context) {
	c.JSON(200,"world")
}
