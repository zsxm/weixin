//scgen
package action

import (
	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/media/index", index)
}

//gen
func index(c chttp.Context) {
	page := c.Param("page")
	action := c.Param("action")
	saveType := c.Param("saveType")
	r := c.NewResult()
	r.Data = map[string]string{
		"action":   action,
		"saveType": saveType,
	}

	c.HTML("/media/"+page, r)
}
