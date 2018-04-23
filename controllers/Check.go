package controllers

import (
	"fmt"
)

type CheckController struct {
	BaseController
}

func(self *CheckController) CheckIp() {
	ip := self.GetString("ip")//获取ip
	port := self.GetString("port")//获取端口号
	proxyType := self.GetString("type")//获取类型

	if  ip == "" || port == "" {
		self.WrongReturn("参数错误",MSG_ERR)
	}

	ip = ip+":"+port
	data,err := self.checkProxy(ip,proxyType)
	if err != nil {
		fmt.Println(err)
		self.WrongReturn(err.Error(),MSG_ERR)
	}


	//p,err := proxyclient.NewProxyClient("socks5://171.115.237.128:57839")


	speed :=data["spend"].(string)

	self.OkReturn("【代理连接成功】：连接时长为："+speed,MSG_OK,data)

}

