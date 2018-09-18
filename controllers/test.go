package controllers

type TestController struct {
	BaseController
}

// @router / [get]
func (this *TestController) SayHello() {
	this.ResponseSuccess("test", "hello world")
}
