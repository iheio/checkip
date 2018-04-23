package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gamexg/proxyclient"
	"io"
	"time"
	"io/ioutil"
	"fmt"
	"github.com/pkg/errors"
	"strings"
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


func(self *BaseController) checkProxy(ip2port string,proxyType string)(map[string]interface{},error) {

	var address string
	switch proxyType {
	case "http":
		address = "http://"+ip2port
		break
	case "https":
		address = "https://"+ip2port
		break
	case "socks5":
		address = "socks5://"+ip2port
		break
	default:
		address = "http://"+ip2port
		break


	}

	p,err := proxyclient.NewProxyClient(address)
	if err != nil {
		panic(err)
	}

	c, err := p.Dial("tcp", "www.baidu.com:80")

	if err != nil {
		fmt.Println(err.Error())
		return nil,errors.New("【代理连接失败】:创建"+proxyType+"隧道失败")
	}


	io.WriteString(c, "GET / HTTP/1.0\r\nHOST:baidu.com\r\n\r\n")
	//计算响应时长
	t1 := time.Now() // get current time
	b, err := ioutil.ReadAll(c)

	if err != nil {
		panic(err)
		return nil,errors.New("【代理连接失败】:无法连接到代理服务器")
	}

	fmt.Println(string(b))

	res := strings.Contains(string(b),"无效用户")
	if res == true {
		return nil,errors.New("【代理连接失败】:此服务器需要用户认证")
	}

	elapsed := time.Since(t1)

	data := make(map[string]interface{})

	fmt.Println("App elapsed: ", elapsed)
	data["spend"] = elapsed.String()

	return data,nil

}