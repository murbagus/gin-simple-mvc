package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/murbagus/gin-simple-mvc/app/models"
)

type PenggunaController struct {
	BaseController
}

// Mendefinisikan routing controller
func (controller *PenggunaController) SetRoute(router *gin.RouterGroup) {
	router.GET("/", controller.Index)
}

func (controller *PenggunaController) Index(c *gin.Context) {
	db := controller.Context.ConnectDatabase("default")
	fmt.Print(db)
	var pengguna models.Pengguna

	db.First(&pengguna)

	c.JSON(200, pengguna)
}
