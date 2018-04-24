package routers

import (
	"github.com/astaxie/beego"
	"ip/controllers"
)

func init() {
    ns := beego.NewNamespace("v1",
    	beego.NSRouter("check_ip", &controllers.CheckController{},"*:CheckIp"),
    	beego.NSRouter("ip", &controllers.IpController{},"get:Get"),
    	beego.NSRouter("ips", &controllers.IpController{},"get:All"),
	)

    beego.AddNamespace(ns)
}
