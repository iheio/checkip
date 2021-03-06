package main

import (
	_ "ip/routers"
	"github.com/astaxie/beego"
	"ip/cros"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cros.Allow(&cros.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))


	beego.Run()
}

