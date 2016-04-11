//scgen
package service

import (
	"fmt"
	tokenapi "weixin/source/token/api"
	"weixin/source/util/webservice"

	"github.com/zsxm/scgo/cjson"
)

const (
	//永久素材发布
	//图片（image）、语音（voice）、视频（video）和缩略图（thumb）
	release_media = "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s"
	//删除永久素材
	del_material = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s"
)

func DeleteMedia(mediaId, userid string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(del_material, token)
	data := `{"media_id":"` + mediaId + `"}`
	cjsn := webservice.SendPostJson(wurl, data)
	return cjsn
}

//接口 发布永久素材
//图片（image）、语音（voice）、和缩略图（thumb）
func ReleaseMediaCommon(filename, userid, mtype string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(release_media, token, mtype)
	body := make(map[string][]string)
	body["media"] = []string{"file", filename}
	return webservice.SendMediaUploadJson(body, wurl)
}

//接口 发布永久素材
//视频（video）
func ReleaseMediaVideo(filename, description, userid, mtype string) *cjson.JSON {
	token := tokenapi.GetCacheToken(userid)
	wurl := fmt.Sprintf(release_media, token, mtype)
	body := make(map[string][]string)
	body["media"] = []string{"file", filename}
	body["description"] = []string{"field", description}
	return webservice.SendMediaUploadJson(body, wurl)
}
