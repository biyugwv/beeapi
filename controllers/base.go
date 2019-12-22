package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	//"github.com/astaxie/beego/logs"
	"fmt"
	"time"

)



type BaseController struct {
	beego.Controller
}

func CheckSSID (ssid string)  bool {
	bm, err := cache.NewCache("redis", `{"conn": "127.0.0.1:6379","password":"wu2182606"}`)
	if err != nil {
		fmt.Println(err)
	}
	v := bm.Get(ssid)
	if v == nil {
		timeoutDuration := 1000 * time.Second
		bm.Put("test","xxxxxxxx",timeoutDuration)
		return  false
	}
	info := string(v.([]byte))
	fmt.Println("info:")
	fmt.Println(info)
	if info == "" {
		return false
	}
	return  true
}