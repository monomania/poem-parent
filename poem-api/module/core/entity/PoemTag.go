package entity

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
)

/**
 * 诗文与标签关系表
 *
 * @author fog
 * @email szy.foggy@gmail.com
 */
type PoemTag struct {
	PoemId int64 `xorm:" comment('诗文id') index"`
	TagId int64 `xorm:" comment('标签id') index"`

	//父类
	entity.Base `xorm:"extends"`
}
