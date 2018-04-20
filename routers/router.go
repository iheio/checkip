package routers

import (
	"github.com/astaxie/beego"
	"ip/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    ns := beego.NewNamespace("v1",
    	beego.NSRouter("check_ip", &controllers.CheckController{},"get:CheckIp"),
	)

    beego.AddNamespace(ns)
}
