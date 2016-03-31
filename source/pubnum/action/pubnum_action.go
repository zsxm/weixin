//scgen
package action

import (
	"weixin/source/pubnum/entity"
	"weixin/source/pubnum/log"
	"weixin/source/pubnum/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	chttp.Action("/pubnum/list", list).Get()
	chttp.Action("/pubnum/add", add).Get()
	chttp.Action("/pubnum/save", save).Post()
	chttp.Action("/pubnum/get/id", get).Get()
}

func get(c chttp.Context) {
	bean := entity.NewPubnumBean()
	e := entity.NewPubnum()
	pri, err := c.Session.Principal()
	if err != nil {
		log.Error(err)
	}
	userid := pri.Id
	e.Userid().SetValue(userid)
	bean.SetEntity(e)
	err = service.PubnumService.Select(bean)
	if err != nil {
		log.Error(err)
	}
	c.JSON(bean.Entitys().JSON(), false)
}

func save(c chttp.Context) {
	e := entity.NewPubnum()
	c.BindData(e)
	pri, err := c.Session.Principal()
	if err != nil {
		log.Error(err)
	}
	userid := pri.Id
	e.Userid().SetValue(userid)             //用户id
	e.Created().SetValue(date.NowUnixStr()) //创建时间
	res, err := service.PubnumService.Save(e)
	if err != nil {
		log.Error(err)
	}
	r, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
	}
	log.Println("add pubnum res row", r)
	c.Redirect("/pubnum/list")
}

func add(c chttp.Context) {
	c.HTML("/pubnum/pubnum", nil)
}

func list(c chttp.Context) {
	bean := entity.NewPubnumBean()
	pri, err := c.Session.Principal()
	if err != nil {
		log.Error(err)
	}
	userid := pri.Id
	e := entity.NewPubnum()
	e.Userid().SetValue(userid)
	e.Userid().FieldExp().Eq().And()
	bean.SetEntity(e)

	service.PubnumService.Select(bean)
	c.HTML("/pubnum/pubnum.list", bean.Entitys())
}
