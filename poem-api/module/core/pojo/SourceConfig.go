package pojo

import "tesou.io/platform/poem-parent/poem-api/module/core/enums"

type SourceConfig struct {
	//数据来源
	S enums.SourceLevel `xorm:" comment('数据来源')"`
	//数据来源对应的id
	SId string `xorm:" comment('数据来源对应的id')"`
	//数据连接
	SUrl string `xorm:" comment('数据连接')"`
	//是否已经同步
	SSynced bool `xorm:" comment('标识是否已经同步')"`
}
