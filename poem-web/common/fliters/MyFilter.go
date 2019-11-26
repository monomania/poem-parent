package fliters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

func init() {
	log = logs.GetBeeLogger()
	beego.InsertFilter("*", beego.BeforeStatic, BeforeStatic, false)
	beego.InsertFilter("*", beego.BeforeRouter, BeforeRouter)
	beego.InsertFilter("*", beego.BeforeExec, BeforeExec)
	beego.InsertFilter("*", beego.AfterExec, AfterExec, false)
	beego.InsertFilter("*", beego.FinishRouter, FinishRouter, false)
}

//BeforeStatic 静态地址之前
func BeforeStatic(ctx *context.Context) {
	log.Debug("BeforeStatic:静态地址之前")
}

//BeforeRouter 寻找路由之前
func BeforeRouter(ctx *context.Context) {
	log.Debug("BeforeRouter:寻找路由之前")
}

//BeforeExec 找到路由之后，开始执行相应的 Controller 之前
func BeforeExec(ctx *context.Context) {
	log.Debug("BeforeExec:找到路由之后，开始执行相应的 Controller 之前")
}

//AfterExec 执行完 Controller 逻辑之后执行的过滤器
func AfterExec(ctx *context.Context) {
	log.Debug("AfterExec:执行完 Controller 逻辑之后执行的过滤器")
}

//FinishRouter 执行完逻辑之后执行的过滤器
func FinishRouter(ctx *context.Context) {
	log.Debug("FinishRouter:执行完逻辑之后执行的过滤器")
}
