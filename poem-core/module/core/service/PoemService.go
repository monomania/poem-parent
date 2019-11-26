package service

import (
	"errors"
	"log"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
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
func (this *PoemService) FindByTitle(title string) (*entity.Poem, error) {
	entitys := make([]*entity.Poem, 0)
	err := mysql.GetEngine().Where(" Title = ?", title).Find(&entitys)
	if err != nil {
		log.Println("FindByTitle:", err)
	}
	if len(entitys) > 0 {
		return entitys[0], nil
	} else {
		return nil, errors.New("FindByTitle:未获取到数据")
	}
}