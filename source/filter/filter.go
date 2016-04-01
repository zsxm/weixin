package filter

import (
	"errors"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	chttp.Add("/*", LoginFilter, "/", "/user/login", "/error404", "/error500", "/weixin/demo")
}

func LoginFilter(fc chttp.FilterContext) error {
	if !fc.Session().IsLogin() {
		fc.HTML("/login", "未登录")
		return errors.New("未登录")
	}
	return nil
}
