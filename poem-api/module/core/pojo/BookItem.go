package pojo

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/pojo"
)

/**
古籍项
*/
type BookItem struct {
	//对应的bookid
	BookId int64 `xorm:" comment('古籍Id') index unique(BookId_TypeName_Name)"`

	TypeName string `xorm:" comment('古籍项类型名称') index unique(BookId_TypeName_Name)"`

	//目录名称
	Name string `xorm:" comment('古籍项名称') index unique(BookId_TypeName_Name) "`

	/**
 	* 简介
 	*/
	Content string `xorm:"mediumtext comment('古籍项内容') "`

	//数据来源
	SourceConfig `xorm:"extends"`

	//父类
	pojo.BasePojo `xorm:"extends"`

}
