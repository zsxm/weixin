//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "common",
	Title:   "common",
	Comment: "common",
}
var control = chttp.NewControl(controlConf)
