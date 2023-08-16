package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"httproxy/controllers"
)

func init() {
	beego.Router("/get_info", &controllers.MainController{}, "get:GetInfo")

}
