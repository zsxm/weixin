//scgen
package api

import (
	"fmt"
	"weixin/source/pubnum/entity"
	"weixin/source/pubnum/log"
	"weixin/source/pubnum/service"

	"github.com/zsxm/scgo/data/cache"
)

const (
	cache_pubnum_key           = "%s_pubnum"
	cache_pubnum_field_appid   = "_appid"
	cache_pubnum_field_secret  = "_secret"
	cache_pubnum_field_wxtoken = "_wxtoken"
	cache_pubnum_field_token   = "_token"
)

func GetPubnum(id string) *entity.Pubnum {
	pubnum := entity.NewPubnum()
	pubnum.Id().SetValue(id)
	pubnum.Id().FieldExp().Eq().And()
	service.PubnumService.SelectOne(pubnum)
	return pubnum
}

type CachePubnum struct {
	Appid   string //第三方用户唯一凭证
	Secret  string //第三方用户唯一凭证密钥，即appsecret
	WxToken string //微信号token
	Token   string //会话access_token

}

//通过公众号id获得公众号基本信息
func SetCachePubNum(pubnum string, cpn CachePubnum) error {
	key := fmt.Sprintf(cache_pubnum_key, pubnum)
	log.Info("设置缓存的公众号基本信息 key=", key)
	if err := cache.HSet(key, cache_pubnum_field_appid, cpn.Appid); err != nil {
		return err
	}
	if err := cache.HSet(key, cache_pubnum_field_secret, cpn.Secret); err != nil {
		return err
	}
	if err := cache.HSet(key, cache_pubnum_field_wxtoken, cpn.WxToken); err != nil {
		return err
	}
	if err := cache.HSet(key, cache_pubnum_field_token, cpn.Token); err != nil {
		return err
	}
	err := cache.Expire(key, 7000)
	return err
}

//通过公众号id获得公众号基本信息
func CachePubNum(pubnum string) CachePubnum {
	cpn := CachePubnum{}
	key := fmt.Sprintf(cache_pubnum_key, pubnum)
	log.Info("获取缓存的公众号基本信息 key=", key)
	if appid, err := cache.HGet(key, cache_pubnum_field_appid); err == nil {
		cpn.Appid = appid
	}
	if secret, err := cache.HGet(key, cache_pubnum_field_secret); err == nil {
		cpn.Secret = secret
	}
	if wxtoken, err := cache.HGet(key, cache_pubnum_field_wxtoken); err == nil {
		cpn.WxToken = wxtoken
	}
	if token, err := cache.HGet(key, cache_pubnum_field_token); err == nil {
		cpn.Token = token
	}
	err := cache.Expire(key, 7000)
	if err != nil {
		log.Error(err)
	}
	return cpn
}