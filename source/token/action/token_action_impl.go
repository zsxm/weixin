//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "token",
	Title:   "Token管理",
	Comment: "Token",
}
var control = chttp.NewControl(controlConf)
