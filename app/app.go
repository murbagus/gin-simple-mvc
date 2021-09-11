package app

import (
	"github.com/gin-gonic/gin"
	"github.com/murbagus/gin-simple-mvc/app/appcontext"
	"github.com/murbagus/gin-simple-mvc/app/middleware"
	"github.com/murbagus/gin-simple-mvc/app/routes"
)

// NewApp membuat app gin baru
// mendaftarkan middleware global dan route ke dalam app
// untuk mendaftarkan middleware global dan route baru silahkan tambahkan
// di dalam fungsi ini
func NewApp() (app *gin.Engine) {
	app = gin.New()
	context := appcontext.NewContext()

	// Default middleware
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// Registrasi global middleware
	// Registrasikan global middlware baru disini
	app.Use(middleware.AuthGlobalMiddleware())

	// Route app
	routes.CreateRouteApi(app, context)

	return
}
