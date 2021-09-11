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

	controllers.CreateController(route, context, "/ping", &controllers.PingController{})
	controllers.CreateController(route, context, "/pengguna", &controllers.PenggunaController{})
}
