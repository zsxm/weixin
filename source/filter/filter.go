package filter

import (
	"errors"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	chttp.Add("/*", LoginFilter, "/", "/user/login", "/error404", "/error500", "/weixin/demo")
}

func LoginFilter(fc chttp.FilterContext) error {
	dmp, err := fc.Session().GetMap()
	if err != nil {
		fc.HTML("/login", "未登录")
		return errors.New("未登录 IP " + fc.Request().RemoteAddr)
	}
	if dmp.Get("id") == "" {
		fc.HTML("/login", "未登录")
		return errors.New("未登录 IP " + fc.Request().RemoteAddr)
	}
	return nil
}
