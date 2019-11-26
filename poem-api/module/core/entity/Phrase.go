package entity

import "tesou.io/platform/poem-parent/poem-api/common/base/entity"

/**
 * 诗句表
 * @author fog
 * @date 2019/10/14
 */
type Phrase struct {
	/**
	  * 诗文id
	  */
	PoemId int64 `xorm:" comment('诗文id') index"`
	//内容
	Content string `xorm:" comment('诗句内容') index"`
	//译
	Yi string `xorm:"varchar(500) comment('译文')"`
	//注
	Zhu string `xorm:"varchar(500) comment('注释')"`
	//父类
	entity.Base `xorm:"extends"`
}
