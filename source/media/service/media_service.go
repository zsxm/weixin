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
func GetMediaList(userid string) *cjson.JSON {
	var mtypes = []string{"news", "image", "video", "voice"}
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(get_media_list, token)
	var data, mtype string
	fmt.Println(wurl, data, mtype)
	offset, count := 0, 20
	var result *cjson.JSON
	for _, v := range mtypes {
	N:
		data = fmt.Sprintf(`{"type":"%s","offset":"%s","count":"%s"}`, v, offset, count)
		resJson := `{"item":[
	{"media_id":"bgftCYzYydQ7v_dbjYhT4uUhqlQrmuA8t2Wvx1TQM5w","name":"upload\/media\/C6AD7356FFB411E5BCB43010B3A0F15C.jpg","update_time":1460358762,
	"url":"https:\/\/mmbiz.qlogo.cn\/mmbiz\/rLRxm6HPRCLY1xbbQKcSHHNPVvhfNa1svicTBzdYvrdKVtCM3pTglTgcvzaThAJ1JLH8IsZswQ6RycwCFfEOXBQ\/0?wx_fmt=jpeg"},
	{"media_id":"CLNWW7Ft6wXNu9-Qf3FgapK-8gYZfDVFrKpg2CxB_Tg","name":"D:\\\\DOWN\\\\golangicon.jpg","update_time":1460355894,
	"url":"https:\/\/mmbiz.qlogo.cn\/mmbiz\/rLRxm6HPRCLY1xbbQKcSHHNPVvhfNa1svicTBzdYvrdKVtCM3pTglTgcvzaThAJ1JLH8IsZswQ6RycwCFfEOXBQ\/0?wx_fmt=jpeg"},
	{"media_id":"CLNWW7Ft6wXNu9-Qf3FgapK-8gYZfDVFrKpg2CxB_Tg","name":"D:\\\\DOWN\\\\golangicon.jpg","update_time":1460355894,
	"url":"https:\/\/mmbiz.qlogo.cn\/mmbiz\/rLRxm6HPRCLY1xbbQKcSHHNPVvhfNa1svicTBzdYvrdKVtCM3pTglTgcvzaThAJ1JLH8IsZswQ6RycwCFfEOXBQ\/0?wx_fmt=jpeg"},
	{"media_id":"CLNWW7Ft6wXNu9-Qf3FgapK-8gYZfDVFrKpg2CxB_Tg","name":"D:\\\\DOWN\\\\golangicon.jpg","update_time":1460355894,
	"url":"https:\/\/mmbiz.qlogo.cn\/mmbiz\/rLRxm6HPRCLY1xbbQKcSHHNPVvhfNa1svicTBzdYvrdKVtCM3pTglTgcvzaThAJ1JLH8IsZswQ6RycwCFfEOXBQ\/0?wx_fmt=jpeg"}
	],"total_count":30,"item_count":20}`
		result = cjson.JsonToMap(resJson)
		totalCount := result.Get("total_count").Integer()
		itemCount := result.Get("item_count").Integer()
		fmt.Println(v, ":", "totalCount=", totalCount, "itemCount=", itemCount, "offset=", offset, "count=", count)
		//当查询的位置数小于总数并且当前返回的数据等于要查询的数据时 继续查询,否则重置查询位置offset=0查询下一个类型
		if offset+count < totalCount && itemCount == count {
			offset += count
			goto N
		} else {
			offset = 0
		}
	}
	//cjsn := webservice.SendPostJson(wurl, data)
	return result
}
