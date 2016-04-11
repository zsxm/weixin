//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
	"weixin/source/media/entity"
)

func init() {
	control.Add("/media/index", index).Get()
}

//gen
func index(c chttp.Context) {
	e := entity.NewMedia()
	c.BindData(e)
	c.JSON(e.JSON(), true)
}
