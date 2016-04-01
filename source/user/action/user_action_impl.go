//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf = chttp.ControlConfigConfig()
var control = chttp.NewControl()

func init() {
	controlConf.SetProject("weixin")
	controlConf.SetModule("user")
	controlConf.SetTitle("用户管理")
	controlConf.SetComment("用户")
	control.Init(controlConf)
}
