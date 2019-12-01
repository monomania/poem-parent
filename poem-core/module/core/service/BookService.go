package service

import (
	"errors"
	"tesou.io/platform/poem-parent/poem-api/common/base"
	entity2 "tesou.io/platform/poem-parent/poem-api/common/base/pojo"
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
	"tesou.io/platform/poem-parent/poem-api/module/core/vo"
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
)

/**
 * 古籍
 * @author fog
 * @date 2019/10/14
 */
type BookService struct {
	mysql.BaseService
}

/**
sv根据名称查找
*/
func (this *BookService) FindByName(name string) (*pojo.Book, error) {
	entitys := make([]*pojo.Book, 0)
	err := mysql.GetEngine().Where(" name = ?", name).Find(&entitys)
	if err != nil {
		base.Log.Info("FindByName:", err)
	}
	if len(entitys) > 0 {
		return entitys[0], nil
	} else {
		return nil, errors.New("FindByName:未获取到数据")
	}
}

/**
分页查询
*/
func (this *BookService) PageBook(v *vo.BookVO, page *entity2.Page) ([]vo.BookVO, error) {
	sql := "select t.* from poem.t_book t where 1=1 "
	if len(v.Name) > 0 {
		sql += " AND t.Name like '%" + v.Name + "%'"
	}

	dataList := make([]vo.BookVO, 0)
	err := this.PageSql(sql, page, &dataList)
	return dataList, err
}
