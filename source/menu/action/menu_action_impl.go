//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf = chttp.ControlConfigConfig()
var control = chttp.NewControl()

func init() {
	controlConf.SetProject("weixin")
	controlConf.SetModule("menu")
	controlConf.SetTitle("菜单管理")
	controlConf.SetComment("菜单")
	control.Init(controlConf)
}
