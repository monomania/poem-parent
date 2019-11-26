package proc

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"regexp"
	"strconv"
	"strings"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen"
)

type PhraseWellProcesser struct {
	phraseWellService service.PhraseWellService
	bookItemService   service.BookItemService
	poemService       service.PoemService
	//sid 与 诗文id的映射
	sId_poemId_map map[string]int64
	//sid 与 古籍id的映射
	sId_bookId_map map[string]int64
}

func GetPhraseWellProcesser() *PhraseWellProcesser {
	return &PhraseWellProcesser{}
}

func (this *PhraseWellProcesser) Startup() {
	//获取所有的bookItem数据
	bookList := []entity.BookItem{}
	this.bookItemService.FindAll(&bookList)
	this.sId_bookId_map = make(map[string]int64, 0)
	for _, e := range bookList {
		this.sId_bookId_map[e.SId] = e.Id
	}
	//获取所有的poem数据
	poemList := []entity.Poem{}
	this.poemService.FindAll(&poemList)
	this.sId_poemId_map = make(map[string]int64, 0)
	for _, e := range poemList {
		this.sId_poemId_map[e.SId] = e.Id
	}

	newSpider := spider.NewSpider(this, "PhraseWellProcesser")
	for i := 1; i <= gushiwen.GSW_MINGJU_MAX_PAGENUM; i++ {
		fullUrl := strings.Replace(gushiwen.GSW_MINGJU_URL, "${pageNum}", strconv.Itoa(i), 1)
		newSpider = newSpider.AddUrl(fullUrl, "html")
	}
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *PhraseWellProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		log.Println("URL:,", request.Url, p.Errormsg())
		return
	}

	data_list_slice := make([]interface{}, 0)
	p.GetHtmlParser().Find(" div.sons div.cont").Each(func(i int, selection *goquery.Selection) {
		data := new(entity.PhraseWell)

		selection.Children().Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				data.Content = strings.TrimSpace(selection.Text())
			} else if i == 1 {
				//nothing
			} else if i == 2 {
				href, exist := selection.Attr("href")
				if exist {

					var regex_temp = regexp.MustCompile(`_\w+\.aspx.*`)
					url_part := regex_temp.FindString(href)
					word := strings.Split(url_part, ".")[0]
					word = strings.Replace(word, "_", "", 1)

					if strings.Contains(href, "shiwenv") {
						data.OwnType = 1
						data.OwnId = this.sId_poemId_map[word]
					} else {
						data.OwnType = 2
						data.OwnId = this.sId_bookId_map[word]
					}
				}
			}
		})

	})
	//执行入库
	if len(data_list_slice) > 0 {
		this.phraseWellService.SaveList(data_list_slice)
	}

}

func (this *PhraseWellProcesser) Finish() {
	log.Println("名句表抓取解析完成 \r\n")

}
