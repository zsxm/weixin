//scgen
package action

import (
	"weixin/source/media/entity"
	"weixin/source/media/log"
	"weixin/source/media/service"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/media/index", index)
}

//gen
func index(c chttp.Context) {
	page := c.Param("page")
	action := c.Param("action")
	saveType := c.Param("saveType")
	ctype := c.Param("ctype")
	copt := c.Param("copt")
	r := c.NewResult()
	if copt == "list" {
		bean := entity.NewMediaBean()
		e := entity.NewMedia()
		e.SaveType().SetValue(saveType)
		e.SaveType().FieldExp().Eq().And()
		e.SetCtype(ctype)
		e.Ctype().FieldExp().Eq().And()
		e.Created().FieldSort().Desc(1)
		bean.SetEntity(e)
		err := service.MediaService.Select(bean)
		if err != nil {
			log.Error(err)
		}
		c.HTML("/media/"+page, bean.Entitys())
	} else {
		r.Data = map[string]string{
			"action":   action,
			"saveType": saveType,
		}
		c.HTML("/media/"+page, r)
	}

}
