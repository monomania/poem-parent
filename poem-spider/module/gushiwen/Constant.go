package gushiwen

const MODULE_FLAG = "gushiwen"

//主页路径
const GSW_URL = "https://so.gushiwen.org"

//诗人采集地址 p=${页码} 最大值暂定10(因其网站限制最多只能查询10页) c=${朝代}
const GSW_POER_MAX_PAGENUM = 10
const GSW_POER_URL = "https://so.gushiwen.org/authors/default.aspx?p=${pageNum}&c=${dynastyLevel}"

//标签获取url
const GSW_TAG_SHIWEN_URL = "https://so.gushiwen.org/shiwen/"
const GSW_TAG_GUWEN_URL = "https://so.gushiwen.org/guwen/"
const GSW_TAG_MINGJU_URL = "https://so.gushiwen.org/mingju/"

//诗文数据
const GSW_SHIWEN_MAX_PAGENUM = 20
const GSW_SHIWEN_URL  = "https://so.gushiwen.org/shiwen/default_0AA${pageNum}.aspx"
//诗句数据
const GSW_PHRASE_URL  = "https://www.gushiwen.org/nocdn/ajaxshiwencont.aspx?id=${id}&value=yizhushang"

//名句数据
const GSW_MINGJU_MAX_PAGENUM = 20
const GSW_MINGJU_URL  = "https://so.gushiwen.org/mingju/default.aspx?p=${pageNum}&c=&t="

//古籍数据
const GSW_GUWEN_MAX_PAGENUM = 360
const GSW_GUWEN_URL  =  "https://so.gushiwen.org/guwen/book_${pageNum}.aspx"


