package vo

import (
	"container/list"
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
)

/**
诗人
 */
type PoerVO struct {
	pojo.Poer `xorm:"extends"`

	/**
	* 诗人标签列表List<PoerTag>
	 */
	Tags list.List;
}

func (this *PoerVO) TableName() string {
	return "t_poer"
}
