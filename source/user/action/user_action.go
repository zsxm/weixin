//scgen
package action

import (
	"weixin/source/user/entity"
	"weixin/source/user/log"
	"weixin/source/user/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/session"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	chttp.Action("/user/login", login).Post()
}

//gen
func login(c chttp.Context) {
	e := entity.NewUser()
	c.BindData(e)
	res, err := service.UserService.Query("select * from user where account = ? and passwd = ?", e.Account().Value(), e.Passwd().Value())
	if err != nil {
		log.Error(err)
	}
	var id, account, passwd, name string
	if res.Size() > 0 {
		mp := res.Get(0)
		id = mp.Get("id")
		account = mp.Get("account")
		passwd = mp.Get("passwd")
		name = mp.Get("name")
	}
	if account != "" && passwd != "" {
		log.Info("登录成功")
		//跳转公众号管理页面
		princ := session.Principal{
			Id:        id,
			Name:      name,
			LoginName: name,
			LoginTime: date.NowUnixStr(),
		}

		c.Session.SetPrincipal(princ)
		c.Redirect("/pubnum/list")
	} else {
		log.Info("登录失败")
		c.HTML("/login", "登录失败")
	}
}
