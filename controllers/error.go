package controllers

import (
	"fmt"
)

type ErrorController struct {
	BaseController
}

// Abort("404")
func (this *ErrorController) Error404() {
	err := fmt.Errorf("Api not fount. url: %s", this.Ctx.Request.URL.Path)
	this.ResponseHttpError(404, "404", err)
}

// Abort("500")
func (this *ErrorController) Error500() {
	err := fmt.Errorf("Internal server error")
	this.ResponseHttpError(500, "500", err)
}

// Abort("405")
func (this *ErrorController) Error405() {
	err := fmt.Errorf("API not used. url: %s", this.Ctx.Request.URL.Path)
	this.ResponseHttpError(405, "405", err)
}
