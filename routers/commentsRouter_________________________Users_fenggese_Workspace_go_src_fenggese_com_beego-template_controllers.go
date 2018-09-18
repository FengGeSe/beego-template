package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["fenggese.com/beego-template/controllers:TestController"] = append(beego.GlobalControllerRouter["fenggese.com/beego-template/controllers:TestController"],
		beego.ControllerComments{
			Method: "SayHello",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
