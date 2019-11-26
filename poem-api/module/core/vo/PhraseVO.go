package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
)

/**
古籍项
*/
type PhraseVO struct {
	//父类
	entity.Phrase `xorm:"extends"`
}

func (this *PhraseVO) TableName() string {
	return "t_phrase"
}
