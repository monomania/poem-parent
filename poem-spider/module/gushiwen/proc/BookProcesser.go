package proc

import (
	"github.com/PuerkitoBio/goquery"
	"tesou.io/platform/poem-parent/poem-api/common/base"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"regexp"
	"strconv"
	"strings"
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen"
)

type BookProcesser struct {
	bookService     service.BookService
	bookItemService service.BookItemService
	url_entity_map  map[string]string
}

func GetBookProcesser() *BookProcesser {
	return &BookProcesser{}
}

func (this *BookProcesser) Startup() {
	newSpider := spider.NewSpider(this, "BookProcesser")
	this.url_entity_map = make(map[string]string)
	for i := 1; i <= gushiwen.GSW_GUWEN_MAX_PAGENUM; i++ {
		fullUrl := strings.Replace(gushiwen.GSW_GUWEN_URL, "${pageNum}", strconv.Itoa(i), 1)
		newSpider = newSpider.AddUrl(fullUrl, "html")
		this.url_entity_map[fullUrl] = strconv.Itoa(i)
	}
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *BookProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Info("URL:,", request.Url, p.Errormsg())
		return
	}

	//从url中获取朝代信息
	book_list_slice := make(map[*pojo.Book][]*pojo.BookItem, 0)
	p.GetHtmlParser().Find(" div.main3 div.left").Each(func(i int, selection *goquery.Selection) {
		book := new(pojo.Book)
		book.SUrl = request.Url
		book.SId = this.url_entity_map[request.Url]
		/*html, _ := selection.Html()
		base.Log.Info(html)*/
		/*if strings.EqualFold(request.Url,"https://so.gushiwen.org/guwen/book_1.aspx"){
			html, _ := selection.Html()
			base.Log.Info(html)
		}*/
		selection.Find("div.sonspic div.cont").Children().Each(func(i int, selection *goquery.Selection) {
			if i == 1 { //名称
				book.Name = strings.TrimSpace(selection.Text())
			} else if i >= 2 { //简介
				book.Brief += strings.TrimSpace(selection.Text())
			}
		})

		bookItems := make([]*pojo.BookItem, 0)
		selection.Find("div.sons div.bookcont * span a").Each(func(i int, selection *goquery.Selection) {
			item := new(pojo.BookItem)
			type_elem := selection.Parent().Parent().Siblings()
			if type_elem.Is("div.bookMl") {
				item.TypeName = strings.TrimSpace(type_elem.Text())
			}
			//目录名
			item.Name = strings.TrimSpace(selection.Text())
			//通过连接获取内容
			href, exist := selection.Attr("href")
			if exist {
				var regex_temp = regexp.MustCompile(`_\w+\.aspx.*`)
				url_part := regex_temp.FindString(href)
				word := strings.Split(url_part, ".")[0]
				word = strings.Replace(word, "_", "", 1)
				item.SId = word

				if !strings.HasPrefix(href, "http") {
					href = gushiwen.GSW_URL + href
				}
			}
			item.SUrl = href
			//加入
			bookItems = append(bookItems, item)
		})

		if len(book.Name) > 0 {
			val, _ := this.bookService.FindByName(book.Name)
			if val == nil || val.Id <= 0 {
				book_list_slice[book] = bookItems
			}
		}
	})

	//执行入库
	if len(book_list_slice) > 0 {
		for k, v := range book_list_slice {
			this.bookService.Save(k)
			bookItems := make([]interface{}, 0)
			for _, e := range v {
				e.BookId = k.Id
				bookItems = append(bookItems, e)
			}
			this.bookItemService.SaveList(bookItems)
		}

	}
}

func (this *BookProcesser) Finish() {
	base.Log.Info("古文抓取解析完成 \r\n")
	//base.Log.Info("古文项抓取开始 \r\n")

}
