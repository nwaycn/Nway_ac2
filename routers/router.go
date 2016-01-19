package routers

import (
	"github.com/nwaycn/Nway_ac2/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{})
}
