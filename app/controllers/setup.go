package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/murbagus/gin-simple-mvc/app/appcontext"
)

type ControllerInterface interface {
	// SetRoute merupakan fungsi yang mendefinisikan routing setiap fungsi controller
	// fungsi ini dipanggil pada controller yang bersangkutan
	SetRoute(router *gin.RouterGroup)

	SetContext(context *appcontext.Context)
}

// BaseController akan di implementasikan untuk setiap controller
// sehingga controller lain memiliki field pada struct BaseController
type BaseController struct {
	Router  *gin.RouterGroup
	Context *appcontext.Context
}

func (bc *BaseController) SetContext(context *appcontext.Context) {
	bc.Context = context
}

// CreateController membuat controller baru pada route
// fungsi ini juga memberi konteks pada controller yang bersangkutan
func CreateController(router *gin.RouterGroup, context *appcontext.Context, routeName string, controller ControllerInterface) {
	controller.SetContext(context)

	// Definisikan route
	router = router.Group(routeName)
	controller.SetRoute(router)
}
