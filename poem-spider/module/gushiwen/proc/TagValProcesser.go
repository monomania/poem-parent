package proc

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"log"
	"regexp"
	"strings"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
	"tesou.io/platform/poem-parent/poem-core/common/pinyin"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
)

type TagProcesser struct {
	TagValService service.TagValService
	TypeName      string
	TypeCode      string
	SpiderUrl     string
}

func GetTagProcesser() *TagProcesser {
	return &TagProcesser{}
}

func (this *TagProcesser) Startup() {
	newSpider := spider.NewSpider(this, "TagProcesser")
	newSpider = newSpider.AddUrl(this.SpiderUrl, "html")
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *TagProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		log.Println("URL:,", request.Url, p.Errormsg())
		return
	}

	tag_list_slice := make([]interface{}, 0)
	tag_list_update_slice := make([]interface{}, 0)
	var sub_type_name string
	p.GetHtmlParser().Find(" div.left div.titletype div.son2").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		selection.Children().Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				sub_type_name = strings.TrimSpace(selection.Text())
				sub_type_name = strings.Replace(sub_type_name, "：", "", 1)
			} else if (i == 1) {
				selection.Children().Each(func(i int, selection *goquery.Selection) {
					tagV := new(entity.TagVal)
					//1.处理标签的url
					tagUrl, exists := selection.Attr("href")
					if (exists) {
						var regex_temp = regexp.MustCompile(`\/\w+\.aspx.*`)
						url_part := regex_temp.FindString(tagUrl)
						word := strings.Split(url_part, ".")[0]
						tagV.Code = strings.Replace(word, "/", "", 1)

						if !strings.HasPrefix(tagUrl, "http"){
							tagUrl = this.SpiderUrl + url_part[1:]
						}
						tagV.SUrl = strings.TrimSpace(tagUrl)
					}
					tagV.Name = strings.TrimSpace(selection.Text())

					if len(tagV.Name) > 0 {
						tagV.TypeName = this.TypeName + "_" + sub_type_name
						tagV.TypeCode = this.TypeCode + "_" + strings.ToUpper(pinyin.Convert(sub_type_name))

						exist := this.TagValService.Exist(tagV)
						if !exist {
							tag_list_slice = append(tag_list_slice, tagV)
						} else {
							tag_list_update_slice = append(tag_list_update_slice, tagV)
						}
					}
				})
			}
		})
		return true
	})
	//执行入库
	this.TagValService.SaveList(tag_list_slice)
	this.TagValService.ModifyList(tag_list_update_slice)
}

func (this *TagProcesser) Finish() {
	log.Println("标签抓取解析完成 \r\n")

}
