package routers

import (
	"mygoblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.MainController{},"get:LoginPage")
}
