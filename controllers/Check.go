package controllers

import (
	"github.com/gamexg/proxyclient"
	"io"
	"io/ioutil"
	"fmt"
	"time"
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


	var address string
	switch proxyType {
		case "http":
			address = "http://"+ip+":"+port
			break
		case "https":
			address = "https://"+ip+":"+port
			break
		case "socks5":
			address = "socks5://"+ip+":"+port
			break
		default:
			address = "http://"+ip+":"+port
			break

	
	}

	fmt.Println(address)

	//p,err := proxyclient.NewProxyClient("socks5://171.115.237.128:57839")
	p,err := proxyclient.NewProxyClient(address)
	if err != nil {
		panic(err)
	}

	c, err := p.Dial("tcp", "www.baidu.com:80")

	if err != nil {
		self.WrongReturn(err.Error(),MSG_ERR)
	}


	io.WriteString(c, "GET / HTTP/1.0\r\nHOST:baidu.com\r\n\r\n")
	//计算响应时长
	t1 := time.Now() // get current time
	b, err := ioutil.ReadAll(c)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	elapsed := time.Since(t1)

	data := make(map[string]interface{})

	fmt.Println("App elapsed: ", elapsed)
	data["spend"] = elapsed.String()


	self.OkReturn("success",MSG_OK,data)

}

