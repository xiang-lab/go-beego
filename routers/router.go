package routers

import (
	"class/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/abc", &controllers.MainController{})
    beego.Router("/register", &controllers.MainController{})
    /* 当实现了自定义的Get方法，将不会执行默认的Get方法 */
    beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
    beego.Router("/index", &controllers.MainController{}, "get:ShowIndex")
}
