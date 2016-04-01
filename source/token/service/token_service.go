//scgen
package service

import (
	"fmt"
	"weixin/source/webservice"

	"github.com/zsxm/scgo/cjson"
	"github.com/zsxm/scgo/data/cache"
)

const (
	cache_pubnum_session_id_key = "%s_pubnum"
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
func GetWeiXinIP(token string) {
}

//获取缓存里面的公众号id
func GetCachePubNumId(sessionid string) (string, error) {
	key := fmt.Sprintf(cache_pubnum_session_id_key, sessionid)
	pubnumid, err := cache.Get(key)
	if err != nil {
		return "", err
	}
	err = cache.Expire(key, 7000)
	if err != nil {
		return "", err
	}
	return pubnumid, nil
}

//设置缓存里面的公众号id
func SetCachePubNumId(sessionid, pubnumid string) error {
	key := fmt.Sprintf(cache_pubnum_session_id_key, sessionid)
	err := cache.Set(key, pubnumid)
	if err != nil {
		return err
	}
	return nil
}
