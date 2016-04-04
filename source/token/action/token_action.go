//scgen
package action

import (
	"weixin/source/pubnum/api"
	"weixin/source/token/entity"
	"weixin/source/token/log"
	"weixin/source/token/service"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools"
	"github.com/zsxm/scgo/tools/date"
)

func init() {
	control.Add("/token/index", index).Get()
	control.Add("/token/api/get/token", apiGetToken).Get()
	control.Add("/token/api/get/wxip", apiGetWxIP).Get()

}

//gen
func index(c chttp.Context) {
	log.Println("c.Session.Id()", c.Session().Id())
	c.HTML("/token/token", nil)
}

//获取微信IP
//1 获取登录用户id
//2 根据用户id获取公众号id
//3 根据公众号id获取公众号基本信息
//4 根据token获取wxip
func apiGetWxIP(c chttp.Context) {

	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	token := service.GetCacheToken(userid)
	cjson := service.GetWeiXinIP(token)
	c.HTML("/token/wxip", cjson.Data())
}

func apiGetToken(c chttp.Context) {
	pubnumid := c.Param("pubnumid")
	var result = c.NewResult()
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
		var access_token string
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
			result.Data = e.Token().Value()
		} else { //超时
			//调用接口
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
					b = true
					result.Code = code
					result.Codemsg = codemsg
					result.Data = cjson.Get("access_token").String()
				} else {
					result.Code = code
					result.Codemsg = codemsg
				}
			} else {
				result.Code = "-100"
				result.Codemsg = "获得token失败"
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
			//设置缓存
			api.SetCachePubNum(pubnumid, cpn)
		}
	} else {
		result.Code = "-100"
		result.Codemsg = "请选择订阅/公众号"
	}
	c.JSON(result, false)
}
