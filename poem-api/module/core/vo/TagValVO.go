package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
)

/**
标签
*/
type TagValVO struct {
	//父类
	entity.TagVal `xorm:"extends"`
	//配置这个Field不进行字段映射
	SUrl string `xorm:"-"`
}

func (this *TagValVO) TableName() string {
	return "t_tag_val"
}
