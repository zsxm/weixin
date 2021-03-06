//scgen
package action

import (
	"fmt"
	"weixin/source/media/entity"
	"weixin/source/media/log"
	"weixin/source/media/service"
	pubnumApi "weixin/source/pubnum/api"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	control.Add("/media/release/common", releaseCommon).Post() //保存普通素材

	control.Add("/media/sync/line/list", syncLineGetList).Get() //同步线上列表
	control.Add("/media/sync/line/media", syncLineMedia).Get()  //同步线上素材

	control.Add("/media/line/get", lineGetList).Get()        //线上列表
	control.Add("/media/release/voice", releaseVoice).Post() //保存视频素材
	control.Add("/media/toupload", toUpload).Get()           //跳转上传页面
	control.Add("/media/delete", deleteMedia).Get()          //删除素材
	control.Add("/media/upload", upload).Post()              //上传素材
	control.Add("/media/news/save", newsSave).Post()         //图文保存
}

//图文保存
func newsSave(c chttp.Context) {
	responseResult := c.NewResult()
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	result, err, _ := service.ReleaseMediaNews(c, userid)
	if err != nil {
		log.Error(err)
	}
	if result.Get("code").String() == "0" { //图文消息保存成功、向数据库保存图文数据
		pubnumId := pubnumApi.GetCachePubNumId(userid)
		mediaId := result.Get("media_id").String()
		e := entity.NewMedia()
		e.SetMediaId(mediaId)
		e.SetCtype("news")
		e.SetPubnumId(pubnumId)
		e.SetSaveType(1)
		e.Created().SetValue(date.NowUnixStr())
		_, err := service.MediaService.Save(e)
		if err != nil {
			log.Error(err)
			responseResult.Code = "120"
			responseResult.Codemsg = "向数据库保存图文消息失败"
		} /*else {
			r, _ := res.RowsAffected()
			if r > 0 {
				for _, tmap := range data {
					tmap["id"] = tools.GEN_SQL_IDENTIFY
					sql, args := tools.MapToSQL_Insert(tmap, "media_news")
					_, err = service.MediaService.Execute(sql, args...)
					if err != nil {
						log.Error(err)
						responseResult.Code = "120"
						responseResult.Codemsg = "向数据库保存图文消息失败"
						break
					}
				}
			}
		}*/
	} else {
		responseResult.Code = result.Get("code").String()
		responseResult.Codemsg = result.Get("codemsg").String()
	}
	c.JSON(responseResult, false)
}

//线上-同步素材
func syncLineMedia(c chttp.Context) {
	var ctype = c.Param("type")
	var id = c.Param("id")
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	service.GetMaterial(c.Session().Id(), userid, ctype, id)
}

//同步线上素材文件列表
func syncLineGetList(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	pubnumId := pubnumApi.GetCachePubNumId(userid)
	r := c.NewResult()
	if pubnumId == "" {
		r.Code = "103"
		r.Codemsg = "请启用一个公众号"
		c.JSON(r, false)
	} else {
		cjsn := service.GetMediaList(c.Session().Id(), userid, pubnumId)
		if cjsn.Get("code").String() == "0" {
			syncCount(c)
		}
		r.Code = cjsn.Get("code").String()
		r.Codemsg = cjsn.Get("codemsg").String()
		c.JSON(r, false)
	}
}

//同步素材数量
func syncCount(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	cjsn := service.GetMediaCounts(c.Session().Id(), userid)
	log.Println("cjsn ", cjsn.DataMap())
	if cjsn.Get("code").String() == "0" {
		err = c.Session().SetKeyMap("media_counts", cjsn.DataMap())
		if err != nil {
			log.Error(err)
		}
	}
	//c.JSON(cjsn.DataMap(), false)
}

//获取永久素材列表
func lineGetList(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	pubnumId := pubnumApi.GetCachePubNumId(userid)
	r := c.NewResult()
	if pubnumId == "" {
		r.Code = "103"
		r.Codemsg = "请启用一个公众号"
	} else {
		ctype := c.Param("mediaType")
		bean := entity.NewMediaBean()
		service.GetMediaByType(ctype, pubnumId, bean)
		r.Data = bean.Entitys().JSON()
	}
	c.JSON(r, false)
}

//永久素材上传跳转
func toUpload(c chttp.Context) {
	c.HTML("/media/media.upload", nil)
}

//素材上传
func upload(c chttp.Context) {
	mfile := c.MultiFile()
	mfile.DirName = "media"
	r := c.NewResult()
	err := mfile.Upload("")
	if err != nil {
		log.Error("upload", err)
		r.Code = "-1"
		r.Codemsg = err.Error()
	}
	r.Data = mfile
	c.JSON(r, false)
}

//永久素材发布
//图片（image）、语音（voice）和缩略图（thumb）
func releaseCommon(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	pubnumId := pubnumApi.GetCachePubNumId(userid)
	e := entity.NewMedia()
	c.BindData(e)
	r := c.NewResult()
	if e.LocalName().Value() == "" {
		r.Code = "101"
		r.Codemsg = "请先上传素材文件"
	} else if pubnumId == "" {
		r.Code = "103"
		r.Codemsg = "请启用一个公众号"
	} else {
		cjsn := service.ReleaseMediaCommon(e.LocalName().Value(), userid, e.Ctype().Value())
		if cjsn != nil {
			r.Code = cjsn.Get("code").String()
			r.Codemsg = cjsn.Get("codemsg").String()
			var mediaId, created, url string
			mediaId = cjsn.Get("media_id").String()
			created = cjsn.Get("created_at").String()
			url = cjsn.Get("url").String()
			if r.Code == "0" {
				r.Data = mediaId
				e.SetMediaId(mediaId)
				if created == "" {
					created = date.NowUnixStr()
				}
				e.Created().SetValue(created)
				e.SetUrl(url)
				e.SetSaveType(1) //永久素材
				e.SetPubnumId(pubnumId)
				_, err = service.MediaService.Save(e)
				if err != nil {
					r.Code = "101"
					r.Codemsg = "素材上成功，但保存数据失败！"
				}
			}
			log.Info(cjsn.Data())
		} else {
			r.Code = "102"
			r.Codemsg = "素材上传失败"
		}
	}
	c.JSON(r, false)
}

//永久素材发布
//视频（video）
func releaseVoice(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	pubnumId := pubnumApi.GetCachePubNumId(userid)
	e := entity.NewMedia()
	c.BindData(e)
	r := c.NewResult()
	if e.LocalName().Value() == "" {
		r.Code = "101"
		r.Codemsg = "请先上传素材文件"
	} else if pubnumId == "" {
		r.Code = "103"
		r.Codemsg = "请启用一个公众号"
	} else {
		//tools.Sleep(5)
		description := `{"title":"%s","introduction":"%s"}`
		description = fmt.Sprintf(description, e.Title().Value(), e.Introduction().Value())
		cjsn := service.ReleaseMediaVideo(e.LocalName().Value(), description, userid, e.Ctype().Value())
		if cjsn != nil {
			r.Code = cjsn.Get("code").String()
			r.Codemsg = cjsn.Get("codemsg").String()
			var mediaId, created, url string
			mediaId = cjsn.Get("media_id").String()
			created = cjsn.Get("created_at").String()
			url = cjsn.Get("url").String()
			if r.Code == "0" {
				e.SetMediaId(mediaId)
				if created == "" {
					created = date.NowUnixStr()
				}
				e.Created().SetValue(created)
				e.SetUrl(url)
				e.SetSaveType(1) //永久素材
				e.SetPubnumId(pubnumId)
				_, err = service.MediaService.Save(e)
				if err != nil {
					r.Code = "101"
					r.Codemsg = "素材上成功，但保存数据失败！"
				}
			}
			//log.Info(cjsn.Data())
		} else {
			r.Code = "102"
			r.Codemsg = "素材上传失败"
		}
	}
	c.JSON(r, false)
}

//删除素材
func deleteMedia(c chttp.Context) {
	mediaId := c.Param("id")
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	cjsn := service.DeleteMedia(c.Session().Id(), mediaId, userid)
	c.JSON(cjsn.Data(), false)
}
