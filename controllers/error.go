package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

// Abort("404")
func (this *ErrorController) Error404() {
	this.Ctx.Output.Status = 404
	this.Data["json"] = map[string]string{"error": "Api not found", "msgs": this.Ctx.Request.URL.Path}
	this.ServeJSON()
}

// Abort("500")
func (this *ErrorController) Error500() {
	this.Ctx.Output.Status = 500
	this.Data["json"] = map[string]string{"error": "Internal server error"}
	this.ServeJSON()
}

// Abort("405")
func (this *ErrorController) Error405() {
	this.Ctx.Output.Status = 405
	this.Data["json"] = map[string]string{"error": "API not used"}
	this.ServeJSON()
}
