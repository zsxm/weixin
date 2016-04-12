//scgen
package action

import (
	"weixin/source/media/entity"
	"weixin/source/media/log"
	"weixin/source/media/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	control.Add("/media/temp/toupload", tempToUpload).Get() //跳转临时素材页面
	control.Add("/media/temp/release", tempRelease).Post()  //临时素材发布
	control.Add("/media/list", tempList).Get()              //临时素材列表
	control.Add("/media/temp/local/get", tempGetList).Get() //本地
}

//获取临时文件列表
func tempGetList(c chttp.Context) {
	mediaType := c.Param("mediaType")
	bean := entity.NewMediaBean()
	media := entity.NewMedia()
	if mediaType != "-1" {
		media.SetCtype(mediaType)
		media.Ctype().FieldExp().Eq().And()
	}
	media.SetSaveType(0)
	media.SaveType().FieldExp().Eq().And()
	media.Created().FieldSort().Desc(1)
	bean.SetEntity(media)
	service.MediaService.Select(bean)
	c.JSON(bean.Entitys().JSON(), false)
}

//临时素材上传跳转
func tempToUpload(c chttp.Context) {
	c.HTML("/media/temp.media.upload", nil)
}

//临时素材发布
func tempRelease(c chttp.Context) {
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
		//添加临时素材
		cjsn := service.AddTempMedia(e.LocalName().Value(), userid, e.Ctype().Value())
		if cjsn != nil {
			r.Code = cjsn.Get("code").String()
			r.Codemsg = cjsn.Get("codemsg").String()
			var mediaId, created, url string
			mediaId = cjsn.Get("media_id").String()
			created = cjsn.Get("created_at").String()
			url = cjsn.Get("url").String()
			if r.Code == "0" {
				e.SetMediaId(mediaId)
				e.Created().SetValue(created)
				if created == "" {
					created = date.NowUnixStr()
				}
				e.SetSaveType(0) //临时素材
				e.SetUrl(url)
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
