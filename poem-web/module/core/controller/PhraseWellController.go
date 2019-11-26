package controller

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-web/common/base/controller"
)

/**
 * 名句
 * @author fog
 * @email szy.foggy@gmail.com
 */
type PhraseWellController struct {
	controller.BaseController
	service.PhraseWellService
}

/**
分页查询
*/
func (this *PhraseWellController) Page() {
	data := &vo.PhraseWellVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.PhraseWellVO, 0)
	err := this.PhraseWellService.Page(data, page, &dataList)
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