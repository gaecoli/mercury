package main

import (
	"mercury/mercury/bootstrap"
	"mercury/mercury/global"
)

func main() {
	bootstrap.InitializeConfig()

	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("logger init success!")

	global.App.DB = bootstrap.InitializeDB()

	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			err := db.Close()
			if err != nil {
				return
			}
		}
	}()

	bootstrap.RunServer()
}
