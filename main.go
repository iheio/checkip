package main

import (
	_ "ip/routers"
	"github.com/astaxie/beego"
	"ip/cros"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cros.Allow(&cros.Options{
					AllowOrigins:     []string{"*"},
					AllowMethods:     []string{"PUT,POST,GET", "PATCH"},
					AllowHeaders:     []string{"Origin"},
					ExposeHeaders:    []string{"Content-Length"},
					AllowCredentials: true,
	}))
	beego.Run()
}

