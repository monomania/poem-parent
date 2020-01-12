package controller

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/pojo"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-core/common/base/controller"
)

/**
 * 诗文与标签关系表
 *
 * @author fog
 * @email szy.foggy@gmail.com
 */
type PoemTagController struct {
	controller.BaseController
	service.PoemTagService
}

/**
分页查询
*/
func (this *PoemTagController) Page() {
	data := &vo.PoemTagVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.PoemTagVO, 0)
	err := this.PoemTagService.Page(data, page, &dataList)
	resp := new(pojo.Response)
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
