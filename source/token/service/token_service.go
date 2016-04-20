//scgen
package service

import (
	"fmt"
	"weixin/source/pubnum/api"
	"weixin/source/token/entity"
	"weixin/source/token/log"
	"weixin/source/util/webservice"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/cjson"
	"github.com/zsxm/scgo/tools"
	"github.com/zsxm/scgo/tools/date"
)

const (
	//获取token
	token_url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	//获取微信服务器IP地址
	weixin_ip = "https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s"
)

//接口 获取token
func GetWeiXinToKen(appid, secret string) *cjson.JSON {
	if appid != "" && secret != "" {
		url := fmt.Sprintf(token_url, appid, secret)
		cjson := webservice.SendGetJson(url, "")
		return cjson
	}
	return nil
}

//接口 获取微信服务器IP地址
func GetWeiXinIP(token string) *cjson.JSON {
	url := fmt.Sprintf(weixin_ip, token)
	cjson := webservice.SendGetJson(url, "")
	return cjson
}

//根据登录用户id，获取缓存中的token
//1 根据用户id获取公众号id
//2 根据公众号id获取公众号基本信息
//3 公众号基本信息获取tokenw
func GetCacheToken(userid string) string {
	pubnumid := api.GetCachePubNumId(userid)
	pubnum := GetPubNum(pubnumid)
	token := pubnum.Token
	return token
}

//根据公众号id获取pubnum
func GetPubNum(pubnumid string) api.CachePubnum {
	return api.CachePubNum(pubnumid)
}

//获取token poor做数据时间比较
//根据公众号获取token并缓存
//1 查询token表验证时间,未超时返回，超时调用接口查询
func GetToKen(pubnumid string, poor bool) chttp.Result {
	var result = chttp.Result{Code: "0", Codemsg: "ok", Data: ""}
	var b bool
	var created string
	if poor {
		res, err := TokenService.Query("select max(created) created from token where pubnumid=?", pubnumid)
		if err != nil {
			log.Error(err)
		}
		if res.Size() > 0 {
			created = res.Get(0).Get("created")
			b = tools.TimePoor(created, 7200)
		}
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
		TokenService.SelectOne(e)

		access_token = e.Token().Value()
		result.Data = e.Token().Value()
	} else { //超时
		//调用接口
		cjson := GetWeiXinToKen(pubNum.Appid().Value(), pubNum.Appsecret().Value())
		if cjson != nil {
			code := cjson.Get("code").String()
			codemsg := cjson.Get("codemsg").String()
			if code == "0" {
				access_token = cjson.Get("access_token").String()
				e.SetToken(access_token)
				e.SetPubnumid(pubnumid)
				e.Created().SetValue(date.NowUnixStr())
				if _, err := TokenService.Save(e); err != nil {
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
			Name:    pubNum.Name().Value(),
			Appid:   pubNum.Appid().Value(),
			Secret:  pubNum.Appsecret().Value(),
			WxToken: pubNum.Token().Value(),
			Token:   access_token,
			Userid:  pubNum.Userid().Value(),
		}
		//设置缓存
		api.SetCachePubNum(pubnumid, cpn)
	}
	return result
}
