package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type TestController struct {
	BaseController
}

// 用于测试服务运行状态的接口
// @router /test [get]
func (this *TestController) SayHello() {
	appname := beego.AppConfig.String("appname")
	this.Set("appname", appname)
	this.ServeJSON()
}

// 用于测试服务器报错信息的接口
// @router /500 [get]
func (this *TestController) Error500() {
	this.ServerError("test:Error500", fmt.Errorf("服务器出错"))
}

// 用于测试客户端请求错误信息的接口
// @router /400 [get]
func (this *TestController) Error400() {
	this.ClientError("test:Error400", fmt.Errorf("客户端请求错误"))
}
