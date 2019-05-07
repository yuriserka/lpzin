package routers

import (
	"github.com/astaxie/beego"
	"github.com/yuriserka/lpzin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
}
