package controllers

import "github.com/astaxie/beego"

type LoginController struct{
	beego.Controller
}
func (this *LoginController) Get(){
	this.TplName="login.tpl"
}
func (this *LoginController) Post(){
	session := this.GetSession("current_user")
	if session !=nil {
		this.TplName="index.tpl"
	}else {
		username := this.GetString("username")
		password := this.GetString("password")

	}
}
