package controllers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
	"github.com/gamexg/proxyclient"
	"io"
	"time"
	"io/ioutil"
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
	chs := make([]chan bool,len(values))
	for k, v := range values {
		ip2port := string(v.([]byte))
		res := strings.Contains(ip2port,":")
		fmt.Println(res)
		if res != true {
			continue
		}
		chs[k] = make(chan bool)
		defer close(chs[k])


		//首先，实现并执行一个匿名的超时等待函数
		timeout := make(chan bool, 1)
		defer close(timeout)
		go func() {
			time.Sleep(time.Duration(3)*time.Second)	//等待1秒钟
			timeout <- true
		}()



	    go this.checkProxy1(ip2port,"http",chs[k])

		//然后，我们把timeout这个channel利用起来
		select {
			case <- timeout:
				chs[k] <- false
			case <-chs[k]:
				ips = append(ips,ip2port)
				if len(ips) == num {//达到了获取数量直接返回
					break

				} else {
					continue
				}

			
		}



		c.Do("HDEL", "useful_proxy",ip2port)






	}

	this.OkReturn("success",MSG_OK,ips)

}

func(this *IpController) checkProxy1(ip2port string,proxyType string,chs chan bool) {

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
		chs <- false
	}


	io.WriteString(c, "GET / HTTP/1.0\r\nHOST:baidu.com\r\n\r\n")
	//计算响应时长
	t1 := time.Now() // get current time
	b, err := ioutil.ReadAll(c)

	if err != nil {
		panic(err)
		chs <- false
	}

	fmt.Println(string(b))

	res := strings.Contains(string(b),"无效用户")
	if res == true {
		chs <- false
	}

	elapsed := time.Since(t1)

	data := make(map[string]interface{})

	fmt.Println("App elapsed: ", elapsed)
	data["spend"] = elapsed.String()

	chs <- true

}