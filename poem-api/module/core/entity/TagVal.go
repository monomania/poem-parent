package entity

import "tesou.io/platform/poem-parent/poem-api/common/base/entity"

type TagVal struct {
	/**
	 * 标签类型
	 */
	TypeName string `xorm:" comment('标签类型名称') index"`
	TypeCode string `xorm:" comment('标签类型编码') index"`
	/**
	 * 标签名称
	 */
	Name string `xorm:" comment('标签名称') index"`
	/**
	标签代码
	*/
	Code string `xorm:" comment('标签编码') index"`

	//数据来源
	SourceConfig `xorm:"extends"`

	//父类
	entity.Base `xorm:"extends"`


}
