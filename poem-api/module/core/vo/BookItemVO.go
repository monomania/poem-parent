package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
)

/**
古籍项
*/
type BookItemVO struct {
	//父类
	entity.BookItem `xorm:"extends"`
	//配置这个Field不进行字段映射
	SUrl string `xorm:"-"`
}

func (this *BookItemVO) TableName() string {
	return "t_book_item"
}
