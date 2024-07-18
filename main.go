package main

import (
	"Raven-Admin/bootstrap"
	g "Raven-Admin/global"
)

func main() {
	bootstrap.InitializeConfig()

	g.Log = bootstrap.InitializeLog()
	g.Log.Info("log init success!")

	g.DB = bootstrap.InitializeDB()
	defer func() {
		if g.DB != nil {
			db, _ := g.DB.DB()
			db.Close()
		}
	}()

	g.Redis = bootstrap.InitializeRedis()

	bootstrap.InitializeValidator()

	bootstrap.RunServer()

}
