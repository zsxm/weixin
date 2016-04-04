//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "pubnum",
	Title:   "公众号管理",
	Comment: "公众号",
}
var control = chttp.NewControl(controlConf)
