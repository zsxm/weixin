package action

import (
	"weixin/source/wx/log"
	"weixin/source/wx/service"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	chttp.Action("/weixin/index", index)
	chttp.Action("/weixin/save", save)
	chttp.Action("/weixin/add", add)
}

func add(c chttp.Context) {
	c.HTML("/weixin/weixin", nil)
}

func save(c chttp.Context) {
	r, err := service.WeixinService.SaveForMap("weixin", c.ParamMaps())
	if err != nil {
		log.Error(err)
	}
	ra, err := r.RowsAffected()
	if err != nil {
		log.Error(err)
	}
	log.Println("save row", ra)
	c.Redirect("/weixin/index")
}

func index(c chttp.Context) {
	log.Println("index")
	c.HTML("/weixin/weixin.list", nil)
}
