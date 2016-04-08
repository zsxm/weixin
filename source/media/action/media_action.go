//scgen
package action

import (
	"weixin/source/media/entity"
	"weixin/source/media/log"
	"weixin/source/media/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/config"
)

func init() {
	control.Add("/media/index", index).Get()
	control.Add("/media/temp/toupload", tempToUpload).Get()
	control.Add("/media/temp/upload", tempUpload).Post()
	control.Add("/media/toupload", toUpload).Get()
	control.Add("/media/upload", upload).Post()
	control.Add("/media/list", tempList).Get()
	control.Add("/media/line/get", tempGetList).Get()       //线上
	control.Add("/media/temp/local/get", tempGetList).Get() //本地
}

//获取临时文件列表
func tempGetList(c chttp.Context) {
	t := c.Param("type")
	mediaType := c.Param("mediaType")
	if t == "local" {
		bean := entity.NewMediaBean()
		media := entity.NewMedia()
		media.SetCtype(mediaType)
		media.Ctype().FieldExp().Eq().And()
		media.Created().FieldSort().Desc(1)
		bean.SetEntity(media)
		service.MediaService.Select(bean)
		c.JSON(bean.Entitys().JSON(), false)
	} else if t == "line" {
		dmp, err := c.Session().GetMap()
		if err != nil {
			log.Error(err)
		}
		userid := dmp.Get("id")
		cjsn := service.GetMediaList(userid, mediaType, "0", "20")
		c.JSON(cjsn.Data(), false)
	}
}

func index(c chttp.Context) {
	e := entity.NewMedia()
	c.BindData(e)
	c.JSON(e.JSON(), true)
}

//临时素材上传跳转
func tempToUpload(c chttp.Context) {
	c.HTML("/media/temp.media.upload", nil)
}

//永久素材上传跳转
func toUpload(c chttp.Context) {
	c.HTML("/media/media.upload", nil)
}

//永久素材上传
func upload(c chttp.Context) {
	mfile := c.MultiFile()
	mfile.DirName = "media"
	err := mfile.Upload("")
	log.Println("error", err)
	r := c.NewResult()
	c.JSON(r, false)
}

//临时素材上传
func tempUpload(c chttp.Context) {
	mfile := c.MultiFile()
	mfile.DirName = "media"
	err := mfile.Upload("")
	r := c.NewResult()

	if err == nil {
		filePath := config.Conf.UploadPath + "/" + mfile.DirName + "/" + mfile.FileNameId[0]
		dmp, err := c.Session().GetMap()
		if err != nil {
			log.Error(err)
		}
		userid := dmp.Get("id")
		cjsn := service.AddTempMedia(filePath, userid, "image")
		if cjsn != nil {
			r.Code = cjsn.Get("code").String()
			r.Codemsg = cjsn.Get("codemsg").String()
			var ctype, mediaId, created string
			ctype = cjsn.Get("type").String()
			mediaId = cjsn.Get("media_id").String()
			created = cjsn.Get("created_at").String()
			e := entity.NewMedia()
			e.SetCtype(ctype)
			e.SetMediaId(mediaId)
			e.SetLocalName(filePath)
			e.Created().SetValue(created)
			_, err = service.MediaService.Save(e)
			if err != nil {
				r.Code = "101"
				r.Codemsg = "素材上成功，但保存数据失败！"
			}
			log.Info(cjsn.Data())
		} else {
			r.Code = "102"
			r.Codemsg = "素材上传失败"
		}
	} else {
		r.Code = "103"
		r.Codemsg = "文件上传失败"
	}
	c.JSON(r, false)
}

//临时素材列表 跳转
func tempList(c chttp.Context) {
	//获取素材数量
	data, err := c.Session().GetKeyMap("media_counts")
	if err != nil {
		log.Error(err)
	}
	if data.Get("code") != "0" {
		dmp, err := c.Session().GetMap()
		if err != nil {
			log.Error(err)
		}
		userid := dmp.Get("id")
		cjsn := service.GetMediaCounts(userid)
		log.Println("cjsn ", cjsn.DataMap())
		if cjsn.Get("code").String() == "0" {
			err = c.Session().SetKeyMap("media_counts", cjsn.DataMap())
			if err != nil {
				log.Error(err)
			}
		}
		c.HTML("/media/temp.media.list", cjsn.Data())
		return
	} else {
		c.HTML("/media/temp.media.list", data)
		return
	}
}
