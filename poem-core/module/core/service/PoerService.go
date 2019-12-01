package service

import (
	"errors"
	"tesou.io/platform/poem-parent/poem-api/common/base"
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
	"tesou.io/platform/poem-parent/poem-api/module/core/enums"
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
)

/**
 * 诗人表
 *
 * @author fog
 * @email szy.foggy@gmail.com
 */
type PoerService struct {
	mysql.BaseService
}


/**
根据朝代名称查找
 */
func (this *PoerService) FindByDynastyAndName(Dynasty enums.DynastyLevel,Name string) (*pojo.Poer, error) {
	entitys := make([]*pojo.Poer, 0)
	err := mysql.GetEngine().Where(" Dynasty = ? AND Name = ? ", Dynasty,Name).Find(&entitys)
	if err != nil {
		base.Log.Info("FindByDynastyAndName:", err)
	}
	if len(entitys) > 0 {
		return entitys[0], nil
	} else {
		return nil, errors.New("FindByDynastyAndName:未获取到数据")
	}
}