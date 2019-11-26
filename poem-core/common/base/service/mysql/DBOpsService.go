package mysql

import (
	"encoding/json"
	"log"
	"tesou.io/platform/poem-parent/poem-api/module/core/entity"
)

type DBOpsService struct {
}

/**
 * 清空表
 */
func (this *DBOpsService) TruncateTable(tables []string) {
	engine := GetEngine()
	for _, v := range tables {
		_, err := engine.Exec(" TRUNCATE TABLE " + v)
		if nil != err {
			log.Println(err)
		}
	}
}

/**
 * xorm支持获取表结构信息，通过调用engine.DBMetas()可以获取到数据库中所有的表，字段，索引的信息。
 */
func (this *DBOpsService) DBMetas() string {
	engine := GetEngine()
	dbMetas, err := engine.DBMetas()
	if nil != err {
		log.Println(err.Error())
	}
	bytes, _ := json.Marshal(dbMetas)
	result := string(bytes)
	return result
}

/**
 * 同步生成数据库表
 */
func (this *DBOpsService) SyncTableStruct() {
	engine := GetEngine()
	var err error
	//sync core model
	err = engine.Sync2(new(entity.Book),new(entity.BookItem), new(entity.Phrase), new(entity.PhraseWell), new(entity.Poem), new(entity.PoemTag), new(entity.Poer), new(entity.PoerTag), new(entity.TagVal))
	if nil != err {
		log.Println(err.Error())
	}
}
