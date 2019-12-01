package service

import (
	"errors"
	"tesou.io/platform/poem-parent/poem-api/common/base"
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
)

type TagValService struct {
	mysql.BaseService
}

/**
sv根据名称查找
 */
func (this *TagValService) FindByName(name string) (*pojo.TagVal, error) {
	entitys := make([]*pojo.TagVal, 0)
	err := mysql.GetEngine().Where(" Name = ?", name).Find(&entitys)
	if err != nil {
		base.Log.Info("FindByName:", err)
	}
	if len(entitys) > 0 {
		return entitys[0], nil
	} else {
		return nil, errors.New("FindByName:未获取到数据")
	}
}
