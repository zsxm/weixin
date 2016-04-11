//scgen
package action

import (
	"fmt"
	"weixin/source/media/entity"
	"weixin/source/media/log"
	"weixin/source/media/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	control.Add("/media/release/common", releaseCommon).Post()
	control.Add("/media/release/voice", releaseVoice).Post()
	control.Add("/media/toupload", toUpload).Get()
	control.Add("/media/delete", deleteMedia).Get()
	control.Add("/media/upload", upload).Post()
}

//永久素材上传跳转
func toUpload(c chttp.Context) {
	c.HTML("/media/media.upload", nil)
}

//永久素材上传
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
	e := entity.NewMedia()
	c.BindData(e)
	r := c.NewResult()
	if e.LocalName().Value() == "" {
		r.Code = "101"
		r.Codemsg = "请先上传素材文件"
	} else {
		cjsn := service.ReleaseMediaCommon(e.LocalName().Value(), userid, e.Ctype().Value())
		if cjsn != nil {
			r.Code = cjsn.Get("code").String()
			r.Codemsg = cjsn.Get("codemsg").String()
			var ctype, mediaId, created, url string
			ctype = cjsn.Get("type").String()
			mediaId = cjsn.Get("media_id").String()
			created = cjsn.Get("created_at").String()
			url = cjsn.Get("url").String()
			if r.Code == "0" {
				e := entity.NewMedia()
				e.SetCtype(ctype)
				e.SetMediaId(mediaId)
				if created == "" {
					created = date.NowUnixStr()
				}
				e.Created().SetValue(created)
				e.SetUrl(url)
				e.SetSaveType(1) //永久素材
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
	e := entity.NewMedia()
	c.BindData(e)
	r := c.NewResult()
	if e.LocalName().Value() == "" {
		r.Code = "101"
		r.Codemsg = "请先上传素材文件"
	} else {
		description := `{"title":"%s","introduction":"%s"}`
		description = fmt.Sprintf(description, e.Title().Value(), e.Introduction().Value())
		cjsn := service.ReleaseMediaVideo(e.LocalName().Value(), description, userid, e.Ctype().Value())
		if cjsn != nil {
			r.Code = cjsn.Get("code").String()
			r.Codemsg = cjsn.Get("codemsg").String()
			var ctype, mediaId, created, url string
			ctype = cjsn.Get("type").String()
			mediaId = cjsn.Get("media_id").String()
			created = cjsn.Get("created_at").String()
			url = cjsn.Get("url").String()
			if r.Code == "0" {
				e.SetCtype(ctype)
				e.SetMediaId(mediaId)
				if created == "" {
					created = date.NowUnixStr()
				}
				e.Created().SetValue(created)
				e.SetUrl(url)
				e.SetSaveType(1) //永久素材
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

func deleteMedia(c chttp.Context) {
	mediaId := c.Param("id")
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	cjsn := service.DeleteMedia(mediaId, userid)
	c.JSON(cjsn.Data(), false)
}
