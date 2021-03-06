package pojo

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/pojo"
)

/**
 * 诗人与标签关系表
 * @author fog
 * @date 2019/10/14
 */
type PoerTag struct {
	PoerId      int64 `xorm:" comment('诗人ID') index"`
	TagId int64 `xorm:" comment('标签ID') index"`

	//父类
	pojo.BasePojo `xorm:"extends"`
}
