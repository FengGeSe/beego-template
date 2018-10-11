package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Resp
}

// 消息格式
type Resp struct {
	Code   string                 `json:"code" desc:"代码"`
	Msg    string                 `json:"msg" desc:"描述"`
	MyData map[string]interface{} `json:"data" desc:"数据"`
}

func (this *BaseController) Set(key string, value interface{}) {
	if this.Resp.MyData == nil {
		this.Code = "ok"
		this.Msg = "success"
		this.MyData = make(map[string]interface{})
	}
	this.MyData[key] = value
	this.Data["json"] = this.Resp
}

// 响应错误
func (this *BaseController) ResponseError(code string, err error) {
	this.Resp.Code = code
	this.Resp.Msg = err.Error()
	this.Ctx.Output.JSON(this.Resp, true, true)
	this.StopRun()
}

// 响应错误，带http状态码
func (this *BaseController) ResponseHttpError(status int, code string, err error) {
	this.Ctx.Output.Status = status
	this.ResponseError(code, err)
}

// 服务器报错
func (this *BaseController) ServerError(code string, err error) {
	this.ResponseHttpError(500, code, err)
}

// 客户端请求报错
func (this *BaseController) ClientError(code string, err error) {
	this.ResponseHttpError(400, code, err)
}
