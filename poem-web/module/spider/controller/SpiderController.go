package controller

import "tesou.io/platform/poem-parent/poem-web/common/base/controller"

type SpiderController struct {
	controller.BaseController
}


func (this *SpiderController) full() {
	this.Data["json"] = "spider"
	this.ServeJSON()
}