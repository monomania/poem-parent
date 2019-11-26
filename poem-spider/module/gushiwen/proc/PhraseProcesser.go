package proc

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"regexp"
	"strings"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
	"tesou.io/platform/poem-parent/poem-core/module/core/service"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen"
	"tesou.io/platform/poem-parent/poem-spider/module/gushiwen/down"
)

type PhraseProcesser struct {
	poemService   service.PoemService
	phraseService service.PhraseService
	//sourceid 对应的诗文实体
	sourceId_poem_map map[string]entity.Poem
}

func GetPhraseProcesser() *PhraseProcesser {
	return &PhraseProcesser{}
}

func (this *PhraseProcesser) Startup() {
	newSpider := spider.NewSpider(this, "PhraseProcesser")

	var dataList []entity.Poem
	this.poemService.FindAll(&dataList)
	this.sourceId_poem_map = make(map[string]entity.Poem, len(dataList))
	for _, data := range dataList {
		//无法强制转换使用json处理
		fullUrl := strings.Replace(gushiwen.GSW_PHRASE_URL, "${id}", data.SId, 1)
		newSpider = newSpider.AddUrl(fullUrl, "html")
		this.sourceId_poem_map[data.SId] = data
	}

	newSpider.SetDownloader(down.NewPhraseDownloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *PhraseProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		log.Println("URL:,", request.Url, p.Errormsg())
		return
	}

	//从url中获取诗文id
	var regex_temp = regexp.MustCompile(`id=.*&`)
	sourceId := strings.Split(regex_temp.FindString(request.Url), "=")[1]
	sourceId = strings.TrimSpace(strings.Replace(sourceId, "&", "", 1))

	poem := this.sourceId_poem_map[sourceId]

	data_list_slice := make([]interface{}, 0)
	//赏析开始
	var shang_point bool
	//参考资料开始
	var reference_point bool
	p.GetHtmlParser().Find("body").Children().Each(func(i int, selection *goquery.Selection) {
		data := new(entity.Phrase)
		data.PoemId = poem.Id
		if selection.Is("div.hr") { //到了赏析部分
			shang_point = true
		} else if strings.Contains(selection.Text(), "参考资料") {
			shang_point = false
			reference_point = true
		}else if shang_point {
			poem.Shang += strings.TrimSpace(selection.Text())
		}  else if reference_point {
			poem.Reference += strings.TrimSpace(selection.Text())
		} else if selection.Is("p") && !shang_point && !reference_point {
			ret, _ := selection.Html()
			retArr := strings.Split(ret, "<br/>")
			data.Content = trimHtml(retArr[0])
			if len(retArr) > 1 {
				data.Yi = trimHtml(retArr[1])
			}
			if len(retArr) > 2 {
				data.Zhu = trimHtml(retArr[2])
			}
		}

		if len(data.Content) > 0 && len(data.Yi) > 0 {
			data_list_slice = append(data_list_slice, data)
		}
	})
	//执行入库
	if len(data_list_slice) > 0 {
		this.phraseService.SaveList(data_list_slice)
		this.poemService.Modify(&poem)
	}

}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func (this *PhraseProcesser) Finish() {
	log.Println("诗句抓取解析完成 \r\n")

}
