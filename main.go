package main

import (
	_ "BeeApi/routers"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:wht284604@tcp(127.0.0.1:3306)/test_db?charset=utf8")
	//orm.RunSyncdb("default", false, true)
}
