//scgen
package action

import (
	pubNumapi "weixin/source/pubnum/api"
	"weixin/source/token/log"
	"weixin/source/token/service"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/token/index", index).Get()
	control.Add("/token/api/get/token", apiGetToken).Get()
	control.Add("/token/api/get/wxip", apiGetWxIP).Get()

}

type data struct {
	Id    string
	Token string
}

//gen
func index(c chttp.Context) {
	dmp, err := c.Session().GetMap()
	if err != nil {
		log.Error(err)
	}
	userid := dmp.Get("id")
	pubnumid := pubNumapi.GetCachePubNumId(userid)
	pubnum := pubNumapi.CachePubNum(pubnumid)
	da := data{
		Id:    pubnumid,
		Token: pubnum.Token,
	}
	c.HTML("/token/token", da)
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
	cjson := service.GetWeiXinIP(c.Session().Id(), token)
	c.HTML("/token/wxip", cjson.Data())
}

func apiGetToken(c chttp.Context) {
	pubnumid := c.Param("pubnumid")
	var result = c.NewResult()
	if pubnumid != "-1" {
		result = service.GetToKen(pubnumid, true)
	} else {
		result.Code = "-100"
		result.Codemsg = "请选择订阅/公众号"
	}
	c.JSON(result, false)
}
