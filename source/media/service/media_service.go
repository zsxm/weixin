//scgen
package service

import (
	"fmt"
	"strconv"
	"weixin/source/media/entity"
	"weixin/source/media/log"
	pubnumApi "weixin/source/pubnum/api"
	tokenapi "weixin/source/token/api"
	"weixin/source/util/webservice"

	"github.com/zsxm/scgo/cjson"
	"github.com/zsxm/scgo/data"
	"github.com/zsxm/scgo/data/cache"
)

const (
	//添加临时素材
	add_temp_media = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	//获取临时素材
	get_temp_media = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
	//素材总数
	get_media_counts = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s"
	//同步素材
	get_material = "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s"

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

var mtypes = []string{"news", "image", "video", "voice"}

const (
	media_sync      = "%s_media_sync"
	media_sync_time = "%s_media_sync_time"
)

//接口 获取永久素材列表
func GetMediaList(userid, pubnumId string) *cjson.JSON {
	var result = &cjson.JSON{}
	mstKey := fmt.Sprintf(media_sync_time, pubnumId)
	mst, _ := cache.Get(mstKey)
	if mst == "" {
		cache.Set(mstKey, "3600")
		cache.Expire(mstKey, 3600)
	} else {
		time, _ := cache.TTL(mstKey)
		result.Set("code", "-104")
		result.Set("codemsg", fmt.Sprintf("请%s秒后再同步,一小时只能同步一次。", strconv.Itoa(time)))
		return result
	}
	msKey := fmt.Sprintf(media_sync, userid)
	cms, err := cache.Get(msKey)
	//同步结束去除同步标记
	defer cache.Delete(msKey)
	if err != nil {
		result.Set("code", "-100")
		result.Set("codemsg", "获取同步标记出错")
		return result
	}
	if cms != "" {
		result.Set("code", "-110")
		result.Set("codemsg", "正在同步中...")
		return result
	} else {
		cache.Set(msKey, "sync")
	}
	//mtypes = []string{"image"}
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(get_media_list, token)
	var data string
	fmt.Println(wurl)
	offset, count := 0, 20
	for _, v := range mtypes {
	N:
		data = fmt.Sprintf(`{"type":"%s","offset":%s,"count":%s}`, v, strconv.Itoa(offset), strconv.Itoa(count))

		//cjsn := cjson.JsonToMap(newsresult)
		cjsn := webservice.SendPostJson(wurl, data)
		code := cjsn.Get("code").String()
		if code == "0" {
			totalCount := cjsn.Get("total_count").Integer()
			itemCount := cjsn.Get("item_count").Integer()
			if totalCount > 0 && itemCount > 0 {
				sync(cjsn, v, pubnumId)

				//fmt.Println(v, ":", "totalCount=", totalCount, "itemCount=", itemCount, "offset=", offset, "count=", count)
				//当查询的位置数小于总数并且当前返回的数据等于要查询的数据时 继续查询,否则重置查询位置offset=0查询下一个类型
				if offset+count < totalCount && itemCount == count {
					offset += count
					cjsn = nil
					goto N
				} else {
					offset = 0
					cjsn = nil
				}
			}
		} else {
			return cjsn
		}
	}
	result.Set("code", "0")
	result.Set("codemsg", "ok")
	return result
}

//同步素材列表
func sync(cjsn *cjson.JSON, ctype, pubnumId string) {
	delSql := "delete from media_news where mediaId = ?"
	items := cjsn.Get("item")
	for i := 0; i < items.Size(); i++ {
		item := items.Index(i)
		mediaId := item.Get("media_id").String()
		e := entity.NewMedia()
		e.SetMediaId(mediaId)
		e.MediaId().FieldExp().Eq().And()
		err := MediaService.SelectOne(e)
		if err != nil {
			log.Error(err)
			continue
		}
		updateTime := item.Get("update_time").String()
		url := item.Get("url").String()
		name := item.Get("name").String()
		title := item.Get("title").String()
		introduction := item.Get("introduction").String()
		e.SetSaveType(1)
		e.Created().SetValue(updateTime)
		e.SetMediaId(mediaId)
		e.SetCtype(ctype)
		e.SetUrl(url)
		e.SetLocalName(name)
		e.SetTitle(title)
		e.SetIntroduction(introduction)
		e.SetPubnumId(pubnumId)
		_, err = MediaService.SaveOrUpdate(e)
		if err != nil {
			log.Error(err)
			continue
		}
		if ctype == "news" {
			_, err = MediaService.Execute(delSql, mediaId)
			if err != nil {
				log.Error(err)
				continue
			} else {
				newsItem := item.Get("content").Get("news_item")
				for ni := 0; ni < newsItem.Size(); ni++ {
					nitem := newsItem.Index(ni)
					nitem.Set("mediaId", mediaId)
					_, err = MediaService.SaveForMap("media_news", nitem.DataMaps())
					if err != nil {
						log.Error(err)
						continue
					}
				}
			}
		}
	}
}

//根据素材类型和pubnumId获取素材
func GetMediaByType(ctype, pubnumId string, bean data.EntityBeanInterface) {
	media := entity.NewMedia()
	if ctype != "-1" {
		media.SetCtype(ctype)
		media.Ctype().FieldExp().Eq().And()
	}

	media.SetPubnumId(pubnumId)
	media.PubnumId().FieldExp().Eq().And()
	media.SetSaveType(1)
	media.SaveType().FieldExp().Eq().And()
	media.Created().FieldSort().Desc(1)
	bean.SetEntity(media)
	err := MediaService.Select(bean)
	if err != nil {
		log.Error(err)
	}
}

//同步素材
func GetMaterial(userid, ctype, id string) {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(get_material, token)
	if id == "" { //如果id为空根据类型同步
		if ctype == "-1" { // ctype 不能为空
			return
		} else {
			pubnumId := pubnumApi.GetCachePubNumId(userid)
			if pubnumId == "" { //公众号不能为空
				return
			}
			bean := entity.NewMediaBean()
			GetMediaByType(ctype, pubnumId, bean)
			es := bean.Datas()
			for i := 0; i < es.Len(); i++ {
				e := es.Get(i)
				cjsn := GetMaterialById(e.MediaId().Value(), wurl)
				log.Info(cjsn.Data())
			}
		}
	} else {
		//同步指定id素材
		cjsn := GetMaterialById(id, wurl)
		log.Info(cjsn.Data())
	}
}

//接口 通过 mediaId 获取指定素材
func GetMaterialById(mediaId, wurl string) *cjson.JSON {
	return webservice.SendPostJson(wurl, `{"media_id":"`+mediaId+`"}`)
}

//获取素材 图文模版
func GetNewsAddTemp(pid string) data.QueryResult {
	sql := "select * from media_news_temp where pid = ?"
	res, err := MediaService.Query(sql, pid)
	if err != nil {
		log.Error(err)
	}
	return res
}
