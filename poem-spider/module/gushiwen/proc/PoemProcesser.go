package proc

import (
	"container/list"
	"github.com/PuerkitoBio/goquery"
	"log"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"strconv"
	"strings"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
	"tesou.io/platform/poem-parent/poem-api/module/core/enums"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen"
)

type PoemProcesser struct {
	poemService    service.PoemService
	poerService    service.PoerService
	tagService     service.TagValService
	poemTagService service.PoemService

	SpiderUrl string
}

func GetPoemProcesser() *PoemProcesser {
	return &PoemProcesser{}
}

func (this *PoemProcesser) Startup() {
	newSpider := spider.NewSpider(this, "PoemProcesser")

	for i := 1; i <= gushiwen.GSW_SHIWEN_MAX_PAGENUM; i++ {
		fullUrl := strings.Replace(this.SpiderUrl, "${pageNum}", strconv.Itoa(i), 1)
		newSpider = newSpider.AddUrl(fullUrl, "html")
	}
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *PoemProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		log.Println("URL:,", request.Url, p.Errormsg())
		return
	}

	//从url中获取朝代信息
	poem_list_slice := make(map[*entity.Poem]*list.List, 0)
	poem_list_update_slice := make(map[*entity.Poem]*list.List, 0)
	p.GetHtmlParser().Find(" div.sons").Each(func(i int, selection *goquery.Selection) {
		poem := new(entity.Poem)
		selection.Find("div.cont").Children().Each(func(i int, selection *goquery.Selection) {
			//诗文名称
			if i == 1 {
				poem.Title = strings.TrimSpace(selection.Children().Text())
			} else if selection.HasClass("source") { //作者及朝代
				var level enums.DynastyLevel
				var poerName string
				selection.Children().Each(func(i int, selection *goquery.Selection) {
					if i == 0 { //朝代
						level = enums.GetDynastyLevel(strings.TrimSpace(selection.Text()))
					}
					if i == 2 { //诗人
						poerName = strings.TrimSpace(selection.Text())
					}
				})

				if len(poerName) > 0 {
					var poer *entity.Poer
					var err error
					poer, err = this.poerService.FindByDynastyAndName(level, poerName)
					if nil != err || poer.Id <= 0 {
						poer = new(entity.Poer)
						poer.Dynasty = level
						poer.Name = poerName
						//因这里获取到诗人信息有限，不做更新
						this.poerService.Save(poer)
					}
					poem.PoerId = poer.Id
				}
			} else if selection.HasClass("contson") { //取出诗文的id
				id, exist := selection.Attr("id")
				if exist {
					id = strings.Replace(id, "contson", "", 1)
				}
				poem.S = enums.GUSHIWEN
				poem.SId = id
				//诗文内容
				poem.Content = strings.TrimSpace(selection.Text())
			}
		})

		tag_list_slice := list.New()
		selection.Find("div.tag").Children().Each(func(i int, selection *goquery.Selection) {
			if selection.Is("a") {
				tagName := strings.TrimSpace(selection.Text())
				tag := entity.TagVal{}
				tag.Name = tagName
				tag_list_slice.PushBack(tag)
			}
		})

		if len(poem.Title) > 0 {
			val, _ := this.poemService.FindByTitle(poem.Title)
			if nil == val || val.Id <= 0 {
				poem_list_slice[poem] = tag_list_slice
			} else {
				poem.Id = val.Id
				poem_list_update_slice[poem] = tag_list_slice
			}
		}
	})
	//执行入库
	if len(poem_list_slice) > 0 {
		for k, v := range poem_list_slice {
			this.poemService.Save(k)
			for item := v.Front(); nil != item; item = item.Next() {
				tag := (item.Value).(entity.TagVal)
				//标签不作保存,只作查询，以保存关系
				val, e := this.tagService.FindByName(tag.Name)
				if e == nil {
					poemTag := new(entity.PoemTag)
					poemTag.PoemId = k.Id
					poemTag.TagId = val.Id
					this.poemTagService.Save(poemTag)
				}
			}
		}
	}
	if len(poem_list_update_slice) > 0 {
		for k, _ := range poem_list_update_slice {
			this.poemService.Modify(k)
			//不再处理标签
		}
	}

}

func (this *PoemProcesser) Finish() {
	log.Println("诗文表抓取解析完成 \r\n")

}
