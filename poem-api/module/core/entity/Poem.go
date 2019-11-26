package entity

import (
	"tesou.io/platform/poem-parent/poem-api/common/base/entity"
	"tesou.io/platform/poem-parent/poem-api/module/core/enums"
)

/**
 * 诗文表
 *
 * @author fog
 * @email szy.foggy@gmail.com
 */
type Poem struct {
	//诗人id
	PoerId int64 `xorm:" comment('诗人id') index unique(PoerId_Title)"`
	/**
	诗名
	 */
	Title string `xorm:" comment('诗名') index unique(PoerId_Title)"`
	/**
	只作记录，不作为主要使用
	 */
	Content string `xorm:"varchar(2000) comment('全诗内容，只作记录，不作为主要使用')"`
	//诗词形式
	Type enums.PoemTypeLevel `xorm:" comment('诗词形式') index"`
	//赏
	Shang string `xorm:"text comment('诗文赏析') "`
	//参考资料
	Reference string `xorm:"text comment('参考资料')"`

	//数据来源
	SourceConfig `xorm:"extends"`

	//父类
	entity.Base `xorm:"extends"`

}
