package funmp

import (
	"weixin/source/pubnum/api"

	"github.com/zsxm/scgo/funcmap"
)

//判断公众号是否开启
//用于配置判断使用
func CheckPumNumEnable(userid, pubnumid string) bool {
	cpnm := api.GetCachePubNumId(userid)
	if cpnm == pubnumid {
		return true
	}
	return false
}

func init() {
	funcmap.AddFuncMap("checkPumNumEnable", CheckPumNumEnable)
}
