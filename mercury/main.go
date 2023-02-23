package main

import (
	"github.com/gin-gonic/gin"
	"mercury/mercury/bootstrap"
	"mercury/mercury/global"
	"net/http"
)

func main() {
	bootstrap.InitializeConfig()

	global.App.Log = bootstrap.InitializeLog()
	var logger = global.App.Log

	logger.Info("logger init success!")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run(":" + global.App.Config.App.Port)
}
