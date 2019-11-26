package vo

import (
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
)


type PoerTagVO struct {
	//父类
	entity.PoerTag `xorm:"extends"`
	//配置这个Field不进行字段映射
	SUrl string `xorm:"-"`
}

func (this *PoerTagVO) TableName() string {
	return "t_poer_tag"
}
