//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "sprm",
	Module:  "smallprogram",
	Title:   "微信小程序",
	Comment: "消息",
}
var control = chttp.NewControl(controlConf)
