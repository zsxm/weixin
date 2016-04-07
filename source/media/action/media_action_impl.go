//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "media",
	Title:   "素材管理",
	Comment: "素材",
}
var control = chttp.NewControl(controlConf)
