package action

import (
	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/error404", error404).Get()
	control.Add("/error500", error500).Get()
}

func error404(c chttp.Context) {
	c.HTML("/error404", "404")
}

func error500(c chttp.Context) {
	c.HTML("/error500", "500")
}
