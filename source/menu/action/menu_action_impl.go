//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "menu",
	Title:   "菜单管理",
	Comment: "菜单",
}
var control = chttp.NewControl(controlConf)
