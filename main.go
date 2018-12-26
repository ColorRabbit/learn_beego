package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "hello/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8", beego.AppConfig.String("mysqluser"), beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqldb"))
	orm.RegisterDataBase("default", "mysql", dbLink, 30)
}

func main() {
	beego.Run()
}


