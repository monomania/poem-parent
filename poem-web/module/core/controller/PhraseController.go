package controller

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-web/common/base/controller"
)

/**
 * 诗句
 * @author fog
 * @date 2019/10/14
 */
type PhraseController struct {
	controller.BaseController
	service.PhraseService
}


/**
分页查询
*/
func (this *PhraseController) Page() {
	data := &vo.PhraseVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.PhraseVO, 0)
	err := this.PhraseService.Page(data, page, &dataList)
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