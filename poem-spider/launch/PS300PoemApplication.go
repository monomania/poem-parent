package launch

import (
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen/proc"
)

func Before_poem(){
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_poem","t_poem_tag","t_phrase"})
}

func Spider_poem() {
	processer := proc.GetPoemProcesser()
	processer.SpiderUrl = gushiwen.GSW_SHIWEN_URL
	processer.Startup()
}

