package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
)

/**
古籍项
*/
type PhraseWellVO struct {
	//父类
	entity.PhraseWell `xorm:"extends"`
}

func (this *PhraseWellVO) TableName() string {
	return "t_phrase_well"
}
