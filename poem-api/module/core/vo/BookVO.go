package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
)

/**
古籍
*/
type BookVO struct {
	//父类
	pojo.Book `xorm:"extends"`
	//配置这个Field不进行字段映射
	SUrl string `xorm:"-"`
}

func (this *BookVO) TableName() string {
	return "t_book"
}
