package main

import (
	_ "mygoblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static","static")
	beego.Run()
}

