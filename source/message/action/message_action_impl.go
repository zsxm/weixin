//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "message",
	Title:   "消息管理",
	Comment: "消息",
}
var control = chttp.NewControl(controlConf)
