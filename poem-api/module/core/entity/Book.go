package entity

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
)

/**
古籍
*/
type Book struct {

	Name string `xorm:" comment('古籍名称') index"`

	/**
 	* 简介
 	*/
	Brief string `xorm:"text comment('古籍简介')"`

	//数据来源
	SourceConfig `xorm:"extends"`

	//父类
	entity.Base `xorm:"extends"`

}
