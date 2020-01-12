package controller

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/pojo"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-core/common/base/controller"
)

/**
 * 古籍
 * @author fog
 * @date 2019/10/14
 */
type BookItemController struct {
	controller.BaseController
	service.BookItemService
}

/**
分页查询
*/
func (this *BookItemController) Page() {
	data := &vo.BookItemVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.BookItemVO, 0)
	err := this.BookItemService.Page(data, page, &dataList)
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
