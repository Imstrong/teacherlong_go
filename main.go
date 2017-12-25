package main

import (
	_ "mygoblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//处理静态资源
	beego.SetStaticPath("/static","static")

	beego.Run()
}

