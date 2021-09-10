package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murbagus/gin-simple-mvc/app/appcontext"
	"github.com/murbagus/gin-simple-mvc/app/controllers"
)

func CreateRouteApi(app *gin.Engine, context *appcontext.Context) {
	routePrefix := "/api"

	// Mengambil route setiap controller
	route := app.Group(routePrefix)
	base := &controllers.BaseController{
		Router:  route,
		Context: context,
	}

	controllers.CreateController(base, "/ping", &controllers.PingController{})
}
