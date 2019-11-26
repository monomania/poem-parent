package launch

import (
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen/proc"
)

func Before_phrase(){
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_phrase"})
}

func Spider_phrase() {
	processer := proc.GetPhraseProcesser()
	processer.Startup()
}

