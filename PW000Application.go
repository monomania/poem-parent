package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "tesou.io/platform/poem-parent/poem-core/common/fliters"
	_ "tesou.io/platform/poem-parent/poem-core/common/routers"
)

func main() {
	beego.LoadAppConfig("ini", "conf/app.conf")
	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	//异步输出日志
	logs.Async(1e3)
	//启动
	beego.Run()
}
