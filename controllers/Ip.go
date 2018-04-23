package controllers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
)

type IpController struct {
	BaseController
}


func (this IpController) Get() {
	num,err := this.GetInt("num") //获取数量
	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	values, _ := redis.Values(c.Do("HGETALL", "useful_proxy"))


	var ips []string
	for _, v := range values {
		ip2port := string(v.([]byte))
		res := strings.Contains(ip2port,":")
		fmt.Println(res)
		if res != true {
			continue
		}

		_,err := this.checkProxy(ip2port,"http")

		if err != nil {//删除不可用ip
			c.Do("HDEL", "useful_proxy",ip2port)
			continue
		}

		ips = append(ips,ip2port)


		if len(ips) == num {//达到了获取数量直接返回
			break

		}




	}

	this.OkReturn("success",MSG_OK,ips)

}

