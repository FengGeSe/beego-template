package routers

import (
	"fenggese.com/beego-template/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	// 跨域访问请求
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	// 从配置文件获得domainname和appname
	host := beego.AppConfig.String("domainname") // api.fenggese.com
	appname := beego.AppConfig.String("appname") // beego-template
	namespace := "/" + host + "/" + appname      // /api.fenggese.com/beego-template

	// /host/appname
	ns := beego.NewNamespace(namespace,

		// /host/appname/subjects
		beego.NSNamespace("/test",
			beego.NSNamespace("/v1", beego.NSInclude(&controllers.TestController{})),
		),
	)
	beego.AddNamespace(ns)
}
