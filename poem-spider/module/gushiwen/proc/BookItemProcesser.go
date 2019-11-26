package proc

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"strings"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"unicode/utf8"
)

type BookItemProcesser struct {
	bookItemService service.BookItemService
	//访问路径与bookitem的映射关系
	url_entity_map map[string]entity.BookItem
}

func GetBookItemProcesser() *BookItemProcesser {
	return &BookItemProcesser{}
}

func (this *BookItemProcesser) Startup() {
	dataList := []entity.BookItem{}
	this.bookItemService.FindContextEmpty(&dataList)

	this.url_entity_map = make(map[string]entity.BookItem)
	newSpider := spider.NewSpider(this, "BookItemProcesser")
	for _, e := range dataList {
		//bookItem := (e).(entity.BookItem)
		this.url_entity_map[e.SUrl] = e
		newSpider = newSpider.AddUrl(e.SUrl, "html")
	}
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(80).Run()
}

func (this *BookItemProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		log.Println("URL:,", request.Url, p.Errormsg())
		return
	}

	//通过surl取出对象的bookitem
	bookItem := this.url_entity_map[request.Url]
	p.GetHtmlParser().Find(" div.sons div.cont").Children().Each(func(i int, selection *goquery.Selection) {
		if selection.Is("h1") { //标题
			//bookItem.Name = strings.TrimSpace(selection.Text())
		} else if selection.Is(" div.contson") {
			text := strings.TrimSpace(selection.Text())
			text = FilterEmoji(text)
			bookItem.Content = text
		}
	})
	//执行入库
	this.bookItemService.Modify(&bookItem)

}

func FilterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		}
	}
	return new_content
}

func (this *BookItemProcesser) Finish() {
	log.Println("古籍项抓取解析完成 \r\n")

}
