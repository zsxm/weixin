//scgen
package action

import (
	"weixin/source/pubnum/api"
	"weixin/source/pubnum/entity"
	"weixin/source/pubnum/log"
	"weixin/source/pubnum/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	//chttp.Action("/pubnum/list", list).Get()
	control.Add("/pubnum/list", list).Get()
	chttp.Action("/pubnum/add", add).Get()
	chttp.Action("/pubnum/save", save).Post()
	chttp.Action("/pubnum/get/id", get).Get()
	chttp.Action("/pubnum/enable", enable).Get()
}

//启用公众号
func enable(c chttp.Context) {
	pubnumid := c.Param("pubnumid")
	prin, err := c.Session().Principal()
	if err != nil {
		log.Error(err)
	}
	userid := prin.Id
	result := c.NewResult()
	err = api.SetCachePubNumId(userid, pubnumid)
	if err != nil {
		result.Code = "-1"
		result.Codemsg = "设置缓存里面的公众号pubnumid失败"
	}
	c.JSON(result, false)
}

//根据当前登录用户获取公众号
func get(c chttp.Context) {
	bean := entity.NewPubnumBean()
	e := entity.NewPubnum()
	pri, err := c.Session().Principal()
	if err != nil {
		log.Error(err)
	}
	userid := pri.Id
	e.Userid().SetValue(userid)
	e.Userid().FieldExp().Eq().And()
	bean.SetEntity(e)
	err = service.PubnumService.Select(bean)
	if err != nil {
		log.Error(err)
	}
	c.JSON(bean.Entitys().JSON(), false)
}

//保存
func save(c chttp.Context) {
	e := entity.NewPubnum()
	c.BindData(e)
	pri, err := c.Session().Principal()
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

//跳转添加页面
func add(c chttp.Context) {
	c.HTML("/pubnum/pubnum", nil)
}

//公众号列表
func list(c chttp.Context) {
	bean := entity.NewPubnumBean()
	pri, err := c.Session().Principal()
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
