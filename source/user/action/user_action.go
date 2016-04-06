//scgen
package action

import (
	"weixin/source/user/entity"
	"weixin/source/user/log"
	"weixin/source/user/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	chttp.Action("/", index).Get() //首页登录
	control.Add("/user/login", login).Post()
}

func index(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	if dmp.Get("id") != "" {
		c.Redirect("/pubnum/list")
	} else {
		c.HTML("/login", nil)
	}
}

//gen
func login(c chttp.Context) {
	e := entity.NewUser()
	c.BindData(e)
	res, err := service.UserService.Query("select * from user where account = ? and passwd = ?", e.Account().Value(), e.Passwd().Value())

	if err != nil {
		log.Error(err)
	}
	var id, account, passwd string
	if res.Size() > 0 {
		mp := res.Get(0)
		id = mp.Get("id")
		account = mp.Get("account")
		passwd = mp.Get("passwd")
	}
	if account != "" && passwd != "" {
		log.Info("登录成功")
		//跳转公众号管理页面
		e.SetId(id)
		c.Session().SetEntity(e)
		c.Session().Set("lasttime", date.NowUnixStr())
		c.Redirect("/pubnum/list")
	} else {
		log.Info("登录失败")
		c.HTML("/login", "登录失败")
	}
}
