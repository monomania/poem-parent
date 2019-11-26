package controller

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-web/common/base/controller"
)

type TagValController struct {
	controller.BaseController
	service.TagValService
}


/**
分页查询
*/
func (this *TagValController) Page() {
	data := &vo.TagValVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.TagValVO, 0)
	err := this.TagValService.Page(data, page, &dataList)
	resp := new(entity.Response)
	if nil != err {
		resp.RetCode = -1
		resp.Message = err.Error()
	} else {
		resp.Data = dataList
		resp.Page = page
	}
	this.Data["json"] = resp
	this.ServeJSON()
}