//scgen
package service

import (
	"weixin/source/pubnum/log"
)

//通过pubnum(公众号名称)获得持有的用户id
func GetPubNumUserid(name string) string {
	result, err := PubnumService.Query("select userid from pubnum where name=?", name)
	if err != nil {
		log.Error(err)
		return ""
	}
	return result.Get(0).Get("userid")
}
