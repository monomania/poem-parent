package controller

import "tesou.io/platform/poem-parent/poem-web/common/base/controller"

type IndexController struct {
	controller.BaseController
}

func (this *IndexController) Get() {
	this.Data["json"] = "hello"
	this.ServeJSON()
}
