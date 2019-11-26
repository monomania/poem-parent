package service

import (
	"errors"
	"log"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
)

type TagValService struct {
	mysql.BaseService
}

/**
sv根据名称查找
 */
func (this *TagValService) FindByName(name string) (*entity.TagVal, error) {
	entitys := make([]*entity.TagVal, 0)
	err := mysql.GetEngine().Where(" Name = ?", name).Find(&entitys)
	if err != nil {
		log.Println("FindByName:", err)
	}
	if len(entitys) > 0 {
		return entitys[0], nil
	} else {
		return nil, errors.New("FindByName:未获取到数据")
	}
}
