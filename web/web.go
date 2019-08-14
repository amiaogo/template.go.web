package routers

import (
	"github.com/astaxie/beego"
	"template.go.web/web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
}
