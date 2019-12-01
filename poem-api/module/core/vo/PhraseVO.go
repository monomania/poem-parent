package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
)

/**
古籍项
*/
type PhraseVO struct {
	//父类
	pojo.Phrase `xorm:"extends"`
}

func (this *PhraseVO) TableName() string {
	return "t_phrase"
}
