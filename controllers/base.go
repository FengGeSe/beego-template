package controllers

import (
	"fenggese.com/beego-template/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// 响应错误
func (this *BaseController) ResponseError(err error) {
	this.Ctx.Output.JSON(err, true, true)
	this.StopRun()
}

// 响应成功
func (this *BaseController) ResponseSuccess(key string, value interface{}) {
	response := &models.RespCode{
		Code:    "ok",
		Message: "success",
		Data:    map[string]interface{}{},
	}
	response.Data[key] = value
	this.Ctx.Output.JSON(response, true, true)
	this.StopRun()
}
