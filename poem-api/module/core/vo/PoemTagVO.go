package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
)

/**
古籍项
*/
type PoemTagVO struct {
	//父类
	pojo.PoemTag `xorm:"extends"`
}

func (this *PoemTagVO) TableName() string {
	return "t_poem_tag"
}
