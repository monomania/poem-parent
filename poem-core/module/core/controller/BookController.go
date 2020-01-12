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
type BookController struct {
	controller.BaseController
	service.BookService
}

/**
调用通用的分页查询
*/
func (this *BookController) BasePage() {
	data := &vo.BookVO{}
	page := this.GetPage()

	//仍需处理
	dataList := make([]vo.BookVO, 0)
	err := this.BookService.Page(data, page, &dataList)
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

/**
分页查询
*/
func (this *BookController) Page() {
	data := vo.BookVO{}
	data.Name = this.Input().Get("name")
	page := this.GetPage()

	dataList, err := this.BookService.PageBook(&data, page)
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
