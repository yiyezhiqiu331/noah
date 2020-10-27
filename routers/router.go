package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/welcome", &controllers.WilController{})
	beego.Router("/list", &controllers.ListController{},"*:Get")

	//beego.Router("/menu1", &controllers.Menu1Controller{},"*:Get")

	beego.Router("/public", &controllers.PublicController{})
	beego.Router("/create", &controllers.CreateController{})
	//beego.Router("/login", &controllers.SessionController{})
	//beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleRegister")
	//beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	//beego.Router("/api/v1.0/sessions", &controllers.SessionController{},"post:Login")
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{},"post:Login")
	beego.Router("/api/v1.0/createshell", &controllers.CreateController{})
	beego.Router("/getlist", &controllers.GetListController{},"*:Get")
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/detail", &controllers.DetailController{})
	beego.Router("/ajaxrun", &controllers.DetailController{},"post:AjaxRun")

}
