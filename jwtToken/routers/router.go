package routers

import (
	"jwtToken/filter"
	"jwtToken/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Redirect")
	//路由过滤
    beego.InsertFilter("/api/v1:*GL-:*",beego.BeforeRouter,filter.AuthFilter)
    //获取Token
	beego.Router("/getpass",&controllers.MainController{},"post:GetToken")

	ns:=beego.NewNamespace("/api",
			beego.NSNamespace("/v1",
				//这里只验证成功与否，若是跟项目结合，则换成自己的理由方法即可
				beego.NSRouter("/GL-verify",&controllers.MainController{},"get:VerifyToken"),
				),
		)

	beego.AddNamespace(ns)
}
