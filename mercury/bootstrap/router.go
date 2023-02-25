package bootstrap

import (
	"github.com/gin-gonic/gin"
	"mercury/mercury/global"
	"mercury/mercury/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

func RunServer() {
	r := setupRouter()
	r.Run(":" + global.App.Config.App.Port)
}

