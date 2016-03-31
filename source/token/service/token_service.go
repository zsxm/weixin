//scgen
package service

import (
	"fmt"
	"weixin/source/webservice"

	"github.com/zsxm/scgo/cjson"
)

const (
	token_url  = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	token_url1 = "https://www.baidu.com/?grant_type=client_credential&appid=%s&secret=%s"
)

func GetWeiXinToKen(appid, secret string) *cjson.JSON {
	if appid != "" && secret != "" {
		url := fmt.Sprintf(token_url1, appid, secret)
		cjson := webservice.SendGetJson(url, "")
		return cjson
	}
	return nil
}
