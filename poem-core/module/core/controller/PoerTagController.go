package controller

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/pojo"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-core/common/base/controller"
)

/**
 * 诗人与标签关系表
 * @author fog
 * @date 2019/10/14
 */
type PoerTagController struct {
	controller.BaseController
	service.PoerTagService
}


/**
分页查询
*/
func (this *PoerTagController) Page() {
	data := &vo.PoerTagVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.PoerTagVO, 0)
	err := this.PoerTagService.Page(data, page, &dataList)
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