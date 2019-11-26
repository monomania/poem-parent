package controller

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-web/common/base/controller"
)

/**
 * 诗人表
 *
 * @author fog
 * @email szy.foggy@gmail.com
 */
type PoerController struct {
	controller.BaseController
	service.PoerService
}

/**
分页查询
*/
func (this *PoerController) Page() {
	data := &vo.PoerVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.PoerVO, 0)
	err := this.PoerService.Page(data, page, &dataList)
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
