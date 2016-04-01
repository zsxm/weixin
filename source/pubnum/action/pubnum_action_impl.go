//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf = chttp.ControlConfigConfig()
var control = chttp.NewControl()

func init() {
	controlConf.SetProject("weixin")
	controlConf.SetModule("pubnum")
	controlConf.SetTitle("公众号管理")
	controlConf.SetComment("公众号")
	control.Init(controlConf)
}
