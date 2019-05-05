package routers

import (
	"github.com/astaxie/beego"
	"github.com/proj_LP/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
}
