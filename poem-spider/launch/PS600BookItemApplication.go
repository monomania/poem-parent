package launch

import (
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen/proc"
)

/*func main() {
	Before_book_item()
	Spider_book_item()
}*/

func Before_book_item() {
	//抓取前清空当前比较表
}

func Spider_book_item() {
	processer := proc.GetBookItemProcesser()
	processer.Startup()
}
