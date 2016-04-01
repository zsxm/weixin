//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf = chttp.ControlConfigConfig()
var control = chttp.NewControl()

func init() {
	controlConf.SetProject("weixin")
	controlConf.SetModule("token")
	controlConf.SetTitle("Token管理")
	controlConf.SetComment("Token")
	control.Init(controlConf)
}
