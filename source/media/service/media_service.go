//scgen
package service

import (
	"fmt"
	tokenapi "weixin/source/token/api"
	"weixin/source/util/webservice"

	"github.com/zsxm/scgo/cjson"
)

const (
	//添加临时素材
	add_temp_media = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	//获取临时素材
	get_temp_media = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
	//素材总数
	get_media_counts = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s"
	//永久素材列表
	get_media_list = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s"
)

//接口 添加临时素材
func AddTempMedia(filename, userid, mtype string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(add_temp_media, token, mtype)
	body := make(map[string][]string)
	body["media"] = []string{"file", filename}
	return webservice.SendMediaUploadJson(body, wurl)
}

//接口 获取临时素材
func GetTempMediag(userid, mediaId string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(get_temp_media, token, mediaId)
	return webservice.GetFile(wurl)
}

//接口 素材总数
func GetMediaCounts(userid string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(get_media_counts, token)
	return webservice.SendGetJson(wurl, "")
}

//接口 获取永久素材列表
func GetMediaList(userid, mtype, offset, count string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(get_media_list, token)
	data := fmt.Sprintf(`{"type":"%s","offset":"%s","count":"%s"}`, mtype, offset, count)
	cjsn := webservice.SendPostJson(wurl, data)
	return cjsn
}
