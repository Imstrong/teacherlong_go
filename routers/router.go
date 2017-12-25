package routers

import (
	"mygoblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.MainController{},"get:Index")
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/about",&controllers.MainController{},"get:About")
}
