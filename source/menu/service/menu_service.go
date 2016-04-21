//scgen
package service

import (
	"fmt"
	"weixin/source/menu/log"
	tokenapi "weixin/source/token/api"
	"weixin/source/util/webservice"

	"github.com/zsxm/scgo/cjson"
	"github.com/zsxm/scgo/data"
)

const (
	menu_get    = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s"
	menu_create = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
)

//创建菜单
func CreateMenu(sessionId, userid, data string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	url := fmt.Sprintf(menu_create, token)
	cjsn := webservice.SendPostJson(sessionId, url, data)
	return cjsn
}

//通过接口获得菜单
func GetMenu(sessionId, userid string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	url := fmt.Sprintf(menu_get, token)
	cjsn := webservice.SendGetJson(sessionId, url, "")
	return cjsn
}

//保存菜单类型
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

//获得菜单类型
func MenuTypeList() data.QueryResult {
	var sql = "select * from menu_type"
	res, err := MenuService.Query(sql)
	if err != nil {
		log.Error(err)
	}
	return res
}
