package entity

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
	"tesou.io/platform/poem-parent/poem-api/module/core/enums"
)

/**
 * 诗人表
 *
 * @author fog
 * @email szy.foggy@gmail.com
 */
type Poer struct {
	/**
	* 朝代
	*/
	Dynasty enums.DynastyLevel `xorm:" comment('朝代') index"`
	/**
	 * 姓名
	 */
	Name string `xorm:" comment('名') index"`
	/**
	 * 字
	 */
	NameWord string `xorm:" comment('字') index"`
	/**
	 * 号
	 */
	NameNum string `xorm:" comment('号') index"`

	/**
	 * 简介
	 */
	Brief string `xorm:"text comment('诗人简介')"`

	/**
	头像url
	 */
	 HeadImg string

	//父类
	entity.Base `xorm:"extends"`
}
