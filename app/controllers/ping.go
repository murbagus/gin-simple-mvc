package controllers

import (
	"github.com/gin-gonic/gin"
)

type PingController struct {
	BaseController
}

// Mendefinisikan routing controller
func (controller *PingController) SetRoute(router *gin.RouterGroup) {
	router.GET("/haha", controller.Index)
}

func (controller *PingController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}
