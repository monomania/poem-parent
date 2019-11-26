package service

import (
	"log"
	"tesou.io/platform/poem-parent/poem-core/common/base/service/mysql"
)

/**
 * 古籍
 * @author fog
 * @date 2019/10/14
 */
type BookItemService struct {
	mysql.BaseService
}

/**
查找context为空的bookitems
 */
func (this *BookItemService) FindContextEmpty(entitys interface{}){
	err := mysql.GetEngine().Where(" (Content IS NULL OR  Content = '' ) AND SUrl != '' AND SUrl IS NOT NULL ").Find(entitys)
	if err != nil {
		log.Println("FindContextEmpty:", err)
	}
}
