//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
	"weixin/source/wx/entity"
)

func init() {
	control.Add("/wx/index", index).Get()
}

//gen
func index(c chttp.Context) {
	e := entity.NewMessage()
	c.BindData(e)
	c.JSON(e.JSON(), true)
}
