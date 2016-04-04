//scgen
package action

import (
	"weixin/source/menu/entity"
	"weixin/source/menu/log"
	"weixin/source/menu/service"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/menu/index", index).Get()
	control.Add("/menu/create", create).Get()
	control.Add("/menu/save", save).Post()
	control.Add("/menu/type/list", typeShow).Get()
	control.Add("/menu/type/save", typeSave).Post()
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

//创建
func create(c chttp.Context) {
	c.HTML("/menu/menu", nil)
}

func save(c chttp.Context) {
	e := entity.NewMenu()
	c.BindData(e)
	c.HTML("/menu/menu", nil)
}
