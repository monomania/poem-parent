package service

import (
	"errors"
	"tesou.io/platform/poem-parent/poem-api/common/base"
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
)

/**
 * 诗文
 *
 * @author fog
 * @email szy.foggy@gmail.com
 */
type PoemService struct {
	mysql.BaseService
}


/**
sv根据名称查找
 */
func (this *PoemService) FindByTitle(title string) (*pojo.Poem, error) {
	entitys := make([]*pojo.Poem, 0)
	err := mysql.GetEngine().Where(" Title = ?", title).Find(&entitys)
	if err != nil {
		base.Log.Info("FindByTitle:", err)
	}
	if len(entitys) > 0 {
		return entitys[0], nil
	} else {
		return nil, errors.New("FindByTitle:未获取到数据")
	}
}