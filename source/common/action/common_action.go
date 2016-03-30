package action

import (
	"github.com/zsxm/scgo/chttp"
)

func init() {
	chttp.Action("/error404", error404).Get()
	chttp.Action("/error500", error500).Get()
}

func error404(c chttp.Context) {
	c.HTML("/error404", "404")
}

func error500(c chttp.Context) {
	c.HTML("/error500", "500")
}
