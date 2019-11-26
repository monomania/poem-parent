package vo

import (
	"container/list"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
)

type PoemVO struct {
	entity.Poem `xorm:"extends"`

	/**
	   * 诗文标签列表List<PoemTag>
	   */
	Tags list.List;

	/**
	 * 诗句List<Phrase>
	 */
	Phrases list.List;

	//配置这个Field不进行字段映射
	SUrl string `xorm:"-"`
}
func (this *PoemVO) TableName() string {
	return "t_poem"
}