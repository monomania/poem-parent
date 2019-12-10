package routers

import (
	"github.com/astaxie/beego"
	controller2 "tesou.io/platform/poem-parent/poem-core/module/core/controller"
	"tesou.io/platform/poem-parent/poem-core/module/index/controller"
)

type Routers struct {
}

func init() {
	//首页
	beego.Router("/", &controller.IndexController{})
	//core模块
	beego.AutoRouter(&controller2.BookController{})
	beego.AutoRouter(&controller2.BookItemController{})
	beego.AutoRouter(&controller2.PhraseController{})
	beego.AutoRouter(&controller2.PhraseWellController{})
	beego.AutoRouter(&controller2.PoemController{})
	beego.AutoRouter(&controller2.PoemTagController{})
	beego.AutoRouter(&controller2.PoerController{})
	beego.AutoRouter(&controller2.PoerTagController{})
	beego.AutoRouter(&controller2.TagValController{})

	//spider模块 使用自动路由
	//beego.AutoRouter(&controller3.SpiderController{})

}
