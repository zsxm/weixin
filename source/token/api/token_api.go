//scgen
package api

import (
	"weixin/source/token/service"
)

//根据用户ID获取token
func GetCacheToken(userid string) string {
	return service.GetCacheToken(userid)
}
