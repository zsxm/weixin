//scgen
package action

import (
	"weixin/source/menu/entity"
	"weixin/source/menu/log"
	"weixin/source/menu/service"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	chttp.Action("/menu/index", index).Get()
	chttp.Action("/menu/add", add).Get()
	chttp.Action("/menu/save", save).Post()
	chttp.Action("/menu/type/list", typeShow).Get()
	chttp.Action("/menu/type/save", typeSave).Post()
}

func typeShow(c chttp.Context) {
	data := service.MenuTypeList()
	log.Info(data)
	c.HTML("/menu/menu.type", data.Data)
}

//保存菜单类型
func typeSave(c chttp.Context) {
	service.SaveMenuType(c.ParamMaps())
	c.Redirect("/menu/type/list")
}

//gen
func index(c chttp.Context) {
	c.HTML("/menu/menu.list", nil)
}

func add(c chttp.Context) {
	c.HTML("/menu/menu", nil)
}

func save(c chttp.Context) {
	e := entity.NewMenu()
	c.BindData(e)
	c.HTML("/menu/menu", nil)
}
