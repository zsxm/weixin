//scgen
package service

import (
	"bytes"
	//	"encoding/json"
	"fmt"
	"weixin/source/media/log"
	tokenapi "weixin/source/token/api"
	"weixin/source/util/webservice"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/cjson"
	"github.com/zsxm/scgo/tools"
)

const (
	//添加永久素材
	//图片（image）、语音（voice）、视频（video）和缩略图（thumb）
	add_media = "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s"

	//新增永久图文
	add_news = "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s"
	//图文正文图片上传
	news_upload_media = "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"

	//删除永久素材
	del_material = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s"
)

//删除素材和数据库
func DeleteMedia(mediaId, userid string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(del_material, token)
	data := `{"media_id":"` + mediaId + `"}`
	cjsn := webservice.SendPostJson(wurl, data)
	if cjsn.Get("code").String() == "0" {
		del_sql := "delete from media where mediaId = ?"
		rs, err := MediaService.Execute(del_sql, mediaId)
		if err != nil {
			log.Error(err)
		} else {
			c, err := rs.RowsAffected()
			if c > 0 {
				del_sql = "delete from media_news where mediaId = ?"
				_, err = MediaService.Execute(del_sql, mediaId)
				if err != nil {
					log.Error(err)
				}
			}
		}

	}
	return cjsn
}

//接口 发布永久素材
//图片（image）、语音（voice）、和缩略图（thumb）
func ReleaseMediaCommon(filename, userid, mtype string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(add_media, token, mtype)
	body := make(map[string][]string)
	body["media"] = []string{"file", filename}
	return webservice.SendMediaUploadJson(body, wurl)
}

//接口 发布永久素材
//视频（video）
func ReleaseMediaVideo(filename, description, userid, mtype string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(add_media, token, mtype)
	body := make(map[string][]string)
	body["media"] = []string{"file", filename}
	body["description"] = []string{"field", description}
	return webservice.SendMediaUploadJson(body, wurl)
}

//图文素材
func ReleaseMediaNews(c chttp.Context, userid string) (*cjson.JSON, error, []map[string]string) {
	var result *cjson.JSON
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(add_news, token)

	titles := c.Params("title")                           //标题
	thumb_media_ids := c.Params("thumb_media_id")         //封面图片素材id
	authors := c.Params("author")                         //作者
	digests := c.Params("digest")                         //摘要
	show_cover_pics := c.Params("show_cover_pic")         //是否显示封面
	contents := c.Params("content")                       //正文
	content_source_urls := c.Params("content_source_url") //图文消息的原文地址
	//log.Println("titles=", titles, " thumb_media_ids=", thumb_media_ids, " authors=", authors, " digests=", digests, " show_cover_pics=", show_cover_pics, " contents=", contents, " content_source_urls=", content_source_urls)
	var b bytes.Buffer
	b.WriteString(`{"articles":[`)
	datas := make([]map[string]string, 0, len(titles))
	for i := 0; i < len(titles); i++ {
		if i > 0 {
			b.WriteString(",")
		}
		m := make(map[string]string)
		m["title"] = value(titles, i)
		m["thumb_media_id"] = value(thumb_media_ids, i)
		m["author"] = value(authors, i)
		m["digest"] = value(digests, i)
		pic := value(show_cover_pics, i)
		if pic == "" {
			pic = "0"
		} else {
			pic = "1"
		}
		m["show_cover_pic"] = pic
		//		val, err := json.Marshal(value(contents, i))
		//		if err != nil {
		//			log.Error("json.Marshal error", err)
		//		}
		m["content"] = value(contents, i)
		m["content_source_url"] = value(content_source_urls, i)
		bjsn, err := cjson.MapToJson(m)
		if err != nil {
			result.Set("code", "101")
			result.Set("codemsg", "生成报文出错")
			return result, err, datas
		}
		b.Write(bjsn)
		datas = append(datas, m)
	}
	b.WriteString("]}")

	return webservice.SendPostJson(wurl, b.String()), nil, datas
	//return result, nil
}

func value(values []string, i int) string {
	if tools.Compare(len(values), i) {
		return values[i]
	}
	return ""
}
