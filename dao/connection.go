package dao

import "github.com/astaxie/beego/orm"

func init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("blog","mysql","root:123456@localhost/blog?charset=utf8")
}
