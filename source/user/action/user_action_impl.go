//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "user",
	Title:   "用户管理",
	Comment: "用户",
}
var control = chttp.NewControl(controlConf)
