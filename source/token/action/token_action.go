//scgen
package action

import (
	"fmt"
	"weixin/source/pubnum/api"
	"weixin/source/token/entity"
	"weixin/source/token/log"
	"weixin/source/token/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	chttp.Action("/token/index", index).Get()
	chttp.Action("/token/api/get/token", apiGetToken).Get()
}

//gen
func index(c chttp.Context) {
	log.Println("c.Session.Id()", c.Session.Id())
	c.HTML("/token/token", nil)
}

func apiGetToken(c chttp.Context) {
	pubnumid := c.GetParam("pubnumid")
	if pubnumid != "-1" {
		res, err := service.TokenService.Query("select max(created) created from token where pubnumid=?", pubnumid)
		if err != nil {
			log.Error(err)
		}
		var created string
		var b bool
		if res.Size() > 0 {
			created = res.Get(0).Get("created")
			b = tools.TimePoor(created, 7000)
		}
		var resJson, access_token string
		//获得公众号信息
		pubNum := api.GetPubnum(pubnumid)
		e := entity.NewToken()
		if b { //未超时
			e.Created().SetValue(created)
			e.SetPubnumid(pubnumid)
			e.Created().FieldExp().Eq().And()
			e.Pubnumid().FieldExp().Eq().And()
			service.TokenService.SelectOne(e)
			access_token = e.Token().Value()
			resJson = fmt.Sprintf(`{"code":%s,"codemsg":"%s","token":"%s"}`, "0", "ok", e.Token().Value())

		} else { //超时
			cjson := service.GetWeiXinToKen(pubNum.Appid().Value(), pubNum.Appsecret().Value())
			if cjson != nil {
				code := cjson.Get("code").String()
				codemsg := cjson.Get("codemsg").String()
				if code == "0" {
					access_token = cjson.Get("access_token").String()
					e.SetToken(access_token)
					e.SetPubnumid(pubnumid)
					e.Created().SetValue(date.NowUnixStr())
					if _, err := service.TokenService.Save(e); err != nil {
						log.Error(err)
					}
					resJson = fmt.Sprintf(`{"code":%s,"codemsg":"%s","token":"%s"}`, code, codemsg, cjson.Get("access_token").String())
					b = true
				} else {
					resJson = fmt.Sprintf(`{"code":"%s","codemsg":"%s"}`, code, codemsg)
				}
			} else {
				resJson = `{"code":"-100","codemsg":"获得token失败"}`
			}
		}
		if b {
			//缓存公众号基本信息
			cpn := api.CachePubnum{
				Appid:   pubNum.Appid().Value(),
				Secret:  pubNum.Appsecret().Value(),
				WxToken: pubNum.Token().Value(),
				Token:   access_token,
			}
			api.SetCachePubNum(pubnumid, cpn)
		}
		c.JSON(resJson, false)
	} else {
		c.JSON(`{"code":"-100","codemsg":"请选择订阅/公众号"}`, false)
	}
}

func apiGetWXIP() {
	var token string
	service.GetWeiXinIP(token)
}
