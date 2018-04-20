package controllers

import (
	"github.com/astaxie/beego"
)

const(
	MSG_OK = 0
	MSG_ERR = -1
)


type BaseController struct {
	beego.Controller
}

/**
正确返回json
 */
func ( self *BaseController ) OkReturn(msg interface{},code int,data interface{}) {
	out := make(map[string]interface{})

	out["msg"] = msg
	out["code"] = code
	out["data"] = data

	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}


/**
错误正确返回json
 */
func ( self *BaseController ) WrongReturn(msg interface{},code int) {
	out := make(map[string]interface{})

	out["msg"] = msg
	out["code"] = code

	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

