//scgen
package action

import (
	"fmt"
	"weixin/source/pubnum/api"
	"weixin/source/token/entity"
	"weixin/source/token/log"
	"weixin/source/token/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	chttp.Action("/token/index", index).Get()
	chttp.Action("/token/api/get", apiGet).Get()
}

//gen
func index(c chttp.Context) {
	c.HTML("/token/token", nil)
}

func apiGet(c chttp.Context) {
	pubnumid := c.GetParam("pubnumid")
	if pubnumid != "-1" {
		res, err := service.TokenService.Query("select max(created) created from token where pubnumid=?", pubnumid)
		if err != nil {
			log.Error(err)
		}
		//var created string
		if res.Size() > 0 {
			//	created = res.Get(0).Get("created")
		}
		pubNum := api.GetPubnum(pubnumid)

		cjson := service.GetWeiXinToKen(pubNum.Appid().Value(), pubNum.Appsecret().Value())
		if cjson != nil {
			code := cjson.Get("code").String()
			codemsg := cjson.Get("codemsg").String()
			if code == "0" {
				access_token := cjson.Get("access_token").String()
				e := entity.NewToken()
				e.SetToken(access_token)
				e.SetPubnumid(pubnumid)
				e.Created().SetValue(date.NowUnixStr())
				if _, err := service.TokenService.Save(e); err != nil {
					log.Error(err)
				}
				c.JSON(fmt.Sprintf(`{"code":%s,"codemsg":"%s","token":"%s"}`, code, codemsg, cjson.Get("access_token").String()), false)
			} else {
				c.JSON(fmt.Sprintf(`{"code":"%s","codemsg":"%s"}`, code, codemsg), false)
			}
		} else {
			c.JSON(`{"code":"-100","codemsg":"获得token失败"}`, false)
		}
	} else {
		c.JSON(`{"code":"-100","codemsg":"请选择订阅/公众号"}`, false)
	}
}
