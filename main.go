package main

import (
	"fenggese.com/beego-template/controllers"
	_ "fenggese.com/beego-template/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// 初始化日志
func initLog() {
	logs.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)
	logs.Async()
	logs.SetLogFuncCallDepth(3)
	logs.Info("日志系统初始化完成！")
}

// 初始化错误处理控制器
func initError() {
	beego.ErrorController(&controllers.ErrorController{})
	logs.Info("错误处理控制器注册完成！")
}

// 初始化Redis
func initRedis() {
	logs.Info("Redis初始化完成！")
}

// 初始化Mysql
func initMysql() {
	logs.Info("Mysql初始化完成！")
}

func init() {
	initLog()
	initMysql()
	initRedis()
	initError()
	logs.Info("初始化工作完毕！")
}

func main() {
	beego.Run()
}
