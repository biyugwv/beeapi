package main


import (
	_ "beeapi/routers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "beeapi/im"
)

func init() {
    
    orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlconf := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpwd") + "@" + beego.AppConfig.String("mysqlurl") + "/" + beego.AppConfig.String("mysqldb") + "?charset=utf8"
    err := orm.RegisterDataBase("default", "mysql", mysqlconf,30)
    if err != nil {
	    fmt.Println(err)
    }
}

func main() {
    im.Run()
   
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
	
}
