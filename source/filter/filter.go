package filter

import (
	"errors"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	chttp.Add("/*", LoginFilter, "/", "/user/login", "/error404", "/error500", "/weixin/demo", "/weixin/share")
}

func LoginFilter(fc chttp.FilterContext) error {
	t := fc.Request().Header["X-Requested-With"]
	dmp, err := fc.Session().GetMap()
	if err != nil {
		if len(t) > 0 && t[0] == "XMLHttpRequest" {
			r := fc.NewResult()
			r.Code = "199"
			r.Codemsg = "登录超时或未登录，请重新登录。"
			fc.JSON(r, false)
		} else {
			fc.HTML("/login", "未登录")
		}
		return errors.New("未登录 IP " + fc.Request().RemoteAddr)
	}
	if dmp.Get("id") == "" {
		if len(t) > 0 && t[0] == "XMLHttpRequest" {
			r := fc.NewResult()
			r.Code = "199"
			r.Codemsg = "登录超时或未登录，请重新登录。"
			fc.JSON(r, false)
		} else {
			fc.HTML("/login", "未登录")
		}
		return errors.New("未登录 IP " + fc.Request().RemoteAddr)
	}
	return nil
}
