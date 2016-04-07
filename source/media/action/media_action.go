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
	control.Add("/media/temp/get", tempGet).Get()
}

func index(c chttp.Context) {
	e := entity.NewMedia()
	c.BindData(e)
	c.JSON(e.JSON(), true)
}

//临时素材上传跳转
func tempToUpload(c chttp.Context) {
	c.HTML("/media/media.upload", nil)
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

//临时素材获取
func tempGet(c chttp.Context) {
	c.HTML("/media/media.upload", nil)
}
