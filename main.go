package main

import (
	_ "testBeego/docs"
	_ "testBeego/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true
}