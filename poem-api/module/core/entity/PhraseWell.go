package entity

import "tesou.io/platform/poem-parent/poem-api/common/base/entity"

/**
 * 名句表
 * @author fog
 * @date 2019/10/14
 */
type PhraseWell struct {
	/**
	  * 名句所属的诗文或古籍ID
	  */
	OwnId int64 `xorm:" comment('名句所属的诗文或古籍ID') index unique(OwnId_Content)"`
	/**
	名句所属的诗文或古籍,诗文为1 古籍为2
	 */
	OwnType int8 `xorm:"tinyint comment('名句所属的诗文或古籍,诗文为1 古籍为2') index"`

	//内容
	Content string `xorm:"varchar(500) comment('名句内容') index unique(OwnId_Content)"`

	//父类
	entity.Base `xorm:"extends"`
}
