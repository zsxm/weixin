//scgen
package service

import (
	//	"weixin/source/menu/entity"
	"weixin/source/menu/log"

	"github.com/zsxm/scgo/data"
)

func SaveMenuType(tmp map[string][]string) int {
	tmp["name"] = []string{tmp["value"][0] + "_" + tmp["name"][0]}
	res, err := MenuService.SaveForMap("menu_type", tmp)
	if err != nil {
		log.Error(err)
	}
	r, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
	}
	return int(r)
}

func MenuTypeList() data.QueryResult {
	var sql = "select * from menu_type"
	res, err := MenuService.Query(sql)
	if err != nil {
		log.Error(err)
	}
	return res
}
