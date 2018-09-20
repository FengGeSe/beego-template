package controllers

import (
	"fenggese.com/beego-template/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// 响应错误
func (this *BaseController) ResponseError(code string, err error) {
	response := &models.RespCode{
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	}
	this.Ctx.Output.JSON(response, true, true)
	this.StopRun()
}

// 响应错误，带http状态码
func (this *BaseController) ResponseHttpError(status int, code string, err error) {
	this.Ctx.Output.Status = status
	this.ResponseError(code, err)
}

// 服务器报错
func (this *BaseController) ResponseServerError(code string, err error) {
	this.ResponseHttpError(500, code, err)
}

// 客户端请求报错
func (this *BaseController) ResponseClientError(code string, err error) {
	this.ResponseHttpError(400, code, err)
}

// 响应成功
func (this *BaseController) ResponseSuccess(key string, value interface{}) {
	this.Ctx.Output.Status = 200
	response := &models.RespCode{
		Code:    "ok",
		Message: "success",
		Data:    map[string]interface{}{},
	}
	response.Data[key] = value
	this.Ctx.Output.JSON(response, true, true)
	this.StopRun()
}
