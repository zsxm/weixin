//scgen
package action

import (
	"weixin/source/menu/log"
	"weixin/source/menu/service"
	pubnumapi "weixin/source/pubnum/api"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/menu/get", get).Get()
	control.Add("/menu/create", create).Get()
	control.Add("/menu/save", save).Post()
	control.Add("/menu/delete", deletem).Get()
	control.Add("/menu/type/list", typeList).Get()
	control.Add("/menu/type/save", typeSave).Post()
}

func typeList(c chttp.Context) {
	data := service.MenuTypeList()
	if c.Param("r") == "ajax" {
		c.JSON(data.Data, false)
	} else {
		c.HTML("/menu/menu.type", data.Data)
	}
}

//保存菜单类型
func typeSave(c chttp.Context) {
	service.SaveMenuType(c.ParamMaps())
	c.Redirect("/menu/type/list")
}

//查询菜单
func get(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	cjsn := service.GetMenu(userid)
	c.JSON(cjsn.Data(), false)
}

//创建
func create(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	pubnumid := pubnumapi.GetCachePubNumId(userid)
	pubnum := pubnumapi.GetPubnum(pubnumid)
	c.HTML("/menu/menu", pubnum)
}

//保存菜单
func save(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	data := c.Param("data")
	cjsn := service.CreateMenu(userid, data)
	c.JSON(cjsn.Data(), false)
}

//删除菜单
func deletem(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	log.Println("delete", userid)
}
