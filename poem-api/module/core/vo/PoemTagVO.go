package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
)

/**
古籍项
*/
type PoemTagVO struct {
	//父类
	entity.PoemTag `xorm:"extends"`
}

func (this *PoemTagVO) TableName() string {
	return "t_poem_tag"
}
