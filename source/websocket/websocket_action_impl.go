package websocket

import (
	"github.com/zsxm/scgo/chttp"
)

var controlConf *chttp.ControlConfig = &chttp.ControlConfig{
	Project: "weixin",
	Module:  "websocket",
	Title:   "websocket",
	Comment: "websocket",
}
var control = chttp.NewControl(controlConf)
