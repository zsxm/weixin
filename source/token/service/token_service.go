//scgen
package service

import (
	"fmt"
	"weixin/source/pubnum/api"
	"weixin/source/util/webservice"

	"github.com/zsxm/scgo/cjson"
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
	pubnum := api.CachePubNum(pubnumid)
	token := pubnum.Token
	return token
}
