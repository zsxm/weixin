//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "wx",
	Title:   "微信接口",
	Comment: "微信接口开发",
}
var control = chttp.NewControl(controlConf)
