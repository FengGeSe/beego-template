package controllers

type TestController struct {
	BaseController
}

// 用于测试服务运行状态的接口
// @router / [get]
func (this *TestController) SayHello() {
	this.ResponseSuccess("test", "hello world")
}
