package controller

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-web/common/base/controller"
)

/**
 * 诗文
 *
 * @author fog
 * @email szy.foggy@gmail.com
 */
type PoemController struct {
	controller.BaseController
	service.PoemService
}

/**
分页查询
*/
func (this *PoemController) Page() {
	data := &vo.PoemVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.PoemVO, 0)
	err := this.PoemService.Page(data, page, &dataList)
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
