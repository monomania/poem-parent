package main

import (
	"encoding/json"
	"errors"
	"tesou.io/platform/poem-parent/poem-api/common/base"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"tesou.io/platform/poem-parent/poem-api/module/core/pojo"
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
	service2 "tesou.io/platform/poem-parent/poem-core/module/core/service"
)

func main() {
	//println(pinyin.Convert("世界"))

	//生成数据库表
	genTable()
	//清空数据表
	truncateTable()

	v, _ := GetCurrentPath()

	base.Log.Info(v)

}


func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func test(){
	//测试
	bookService := new(service2.BookService)
	//插入数据
	book := new(pojo.Book)
	save := bookService.Save(book)
	base.Log.Info(save)
	base.Log.Info(book.Id)
	//查询
	dataList := make([]*pojo.Book, 0)
	bookService.FindAll(&dataList)
	bytes, _ := json.Marshal(dataList)
	base.Log.Info("query all : "+string(bytes))
}


func genTable() {
	generateService := new(mysql.DBOpsService)
	generateService.SyncTableStruct()
}

func truncateTable() {
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_book","t_book_item","t_phrase","t_phrase_well","t_poem","t_poem_tag","t_poer","t_poer_tag","t_tag_val"})
}
