package launch

import (
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen/proc"
)

func Before_tag(){
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_tag_val"})
}

func Spider_SHIWEN_TYPE_tag() {
	processer := proc.GetTagProcesser()
	processer.TypeName = "诗文"
	processer.TypeCode = "SHIWEN_TYPE"
	processer.SpiderUrl = gushiwen.GSW_TAG_SHIWEN_URL
	processer.Startup()
}
func Spider_GUWEN_TYPE_tag() {
	processer := proc.GetTagProcesser()
	processer.TypeName = "古文"
	processer.TypeCode = "GUWEN_TYPE"
	processer.SpiderUrl = gushiwen.GSW_TAG_GUWEN_URL
	processer.Startup()
}
func Spider_MINGJU_TYPE_tag() {
	processer := proc.GetTagProcesser()
	processer.TypeName = "名句"
	processer.TypeCode = "MINGJU_TYPE"
	processer.SpiderUrl = gushiwen.GSW_TAG_MINGJU_URL
	processer.Startup()
}
