package proc

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"tesou.io/platform/poem-parent/poem-api/common/base"
	"regexp"
	"strconv"
	"strings"
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
	"tesou.io/platform/poem-parent/poem-api/module/core/enums"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen"
)

type PoerProcesser struct {
	PoerService service.PoerService
}

func GetPoerProcesser() *PoerProcesser {
	return &PoerProcesser{}
}

func (this *PoerProcesser) Startup() {
	newSpider := spider.NewSpider(this, "PoerProcesser")

	for i := 1; i <= gushiwen.GSW_POER_MAX_PAGENUM; i++ {
		halfUrl := strings.Replace(gushiwen.GSW_POER_URL, "${pageNum}", strconv.Itoa(i), 1)
		for _, v := range enums.DynastyMap {
			fullUrl := strings.Replace(halfUrl, "${dynastyLevel}", v, 1)
			newSpider = newSpider.AddUrl(fullUrl, "html")
		}
	}
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *PoerProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Info("URL:,", request.Url, p.Errormsg())
		return
	}

	//从url中获取朝代信息
	var regex_temp = regexp.MustCompile(`c=.*`)
	word := strings.Split(regex_temp.FindString(request.Url), "=")[1]
	level := enums.GetDynastyLevel(word)

	poer_list_slice := make([]interface{}, 0)
	poer_list_update_slice := make([]interface{}, 0)
	p.GetHtmlParser().Find(" div.sonspic").Each(func(i int, selection *goquery.Selection) {
		poer := new(pojo.Poer)
		selection.Find("div.cont").Children().Each(func(i int, selection *goquery.Selection) {
			if i == 0 { //0.头像图片
				headUrl, flag := selection.Find("a img").Attr("src")
				if flag {
					poer.HeadImg = strings.TrimSpace(headUrl)
				}
			} else if i == 1 { //1 名称
				name := selection.Find("a b").Text()
				poer.Name = strings.TrimSpace(name)
			} else if i == 2 { //2简介
				desc := strings.TrimSpace(selection.Text())
				poer.Brief = desc
			} else {
				base.Log.Info(i, selection.Text())
			}
		})
		if len(poer.Name) > 0 {
			//设置朝代
			poer.Dynasty = level

			exist := this.PoerService.Exist(poer)
			if !exist {
				poer_list_slice = append(poer_list_slice, poer)
			} else {
				poer_list_update_slice = append(poer_list_update_slice, poer)
			}
		}
	})
	//执行入库
	this.PoerService.SaveList(poer_list_slice)
	this.PoerService.ModifyList(poer_list_update_slice)

}

func (this *PoerProcesser) Finish() {
	base.Log.Info("诗人表抓取解析完成 \r\n")

}
