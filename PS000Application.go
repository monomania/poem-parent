package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"tesou.io/platform/poem-parent/poem-spider/launch"
)

func main() {
	launch.Before_poer()
	launch.Before_tag()
	launch.Before_poem()
	launch.Before_phrase()
	launch.Before_book()
	launch.Before_book_item()
	launch.Before_phrase_well()
	//诗人数据
	//诗文类型标签
	//诗文数据
	//诗句数据
	//古籍数据
	//古籍项数据
	//名句数据
	launch.Spider_poer()
	launch.Spider_SHIWEN_TYPE_tag()
	launch.Spider_GUWEN_TYPE_tag()
	launch.Spider_MINGJU_TYPE_tag()
	launch.Spider_poem()
	launch.Spider_phrase()
	launch.Spider_book()
	launch.Spider_book_item()
	launch.Spider_phrase_well()

}

func get() {
	//get请求
	//http.Get的参数必须是带http://协议头的完整url,不然请求结果为空

	header := http.Header{}
	header.Add("Host", "www.gushiwen.org")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:70.0) Gecko/20100101 Firefox/70.0")
	header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("DNT", "1")
	header.Add("Connection", "keep-alive")
	//header.Add("Cookie","Hm_lvt_04660099568f561a75456483228a9516=1573978010,1574343926,1574429250,1574555904; Hm_lvt_3c8ecbfa472e76b0340d7a701a04197e=1573978016,1574429256,1574439784; login=flase; Hm_lpvt_04660099568f561a75456483228a9516=1574556334; ASP.NET_SessionId=qveuzvpe5yfwucrf5cjyfxy3")
	header.Add("Upgrade-Insecure-Requests", "1")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	cookies := make([]*http.Cookie, 5)
	cookies[0] = new(http.Cookie)
	cookies[0].Name = "ASP.NET_SessionId"
	cookies[0].Value = "qveuzvpe5yfwucrf5cjyfxy3"
	cookies[1] = new(http.Cookie)
	cookies[1].Name = "Hm_lpvt_04660099568f561a75456483228a9516"
	cookies[1].Value = "1574556334"
	cookies[2] = new(http.Cookie)
	cookies[2].Name = "Hm_lvt_04660099568f561a75456483228a9516"
	cookies[2].Value = "1573978010,1574343926,1574429250,1574555904"
	cookies[3] = new(http.Cookie)
	cookies[3].Name = "Hm_lvt_3c8ecbfa472e76b0340d7a701a04197e"
	cookies[3].Value = "1573978016,1574429256,1574439784"
	cookies[4] = new(http.Cookie)
	cookies[4].Name = "login"
	cookies[4].Value = "flase"

	client := &http.Client{}
	url := "https://www.gushiwen.org/nocdn/ajaxshiwencont.aspx?id=09d31b73b44d&value=yizhushang"
	reqest, err := http.NewRequest("GET", url, nil)

	//reqest.Header = header
	reqest.AddCookie(cookies[0])
	reqest.AddCookie(cookies[1])
	reqest.AddCookie(cookies[2])
	reqest.AddCookie(cookies[3])
	reqest.AddCookie(cookies[4])
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Get request result: %s\n", string(body))
}
