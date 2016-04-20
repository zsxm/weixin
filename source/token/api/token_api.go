//scgen
package api

import (
	"weixin/source/pubnum/api"
	"weixin/source/token/service"
)

//根据用户ID获取当前启用的token
func GetCacheToken(userid string) string {
	return service.GetCacheToken(userid)
}

//根据公众号id获取CachePubNum
func GetCacheTokenByPubnumId(pubnumid string) api.CachePubnum {
	return service.GetPubNum(pubnumid)
}
