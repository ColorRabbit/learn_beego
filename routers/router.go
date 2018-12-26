package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user/?:id:int", &controllers.UserController{})
    beego.Include(&controllers.AdminController{})
}

