package api

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"weixin/source/token/api"
	"weixin/source/util/webservice"
	"weixin/source/wx/log"

	"github.com/zsxm/scgo/data/cache"
	"github.com/zsxm/scgo/tools/date"
	"github.com/zsxm/scgo/tools/uuid"
)

const (
	jsapi_ticket       = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	cache_jsapi_ticket = "%s_jsapi_ticket_pubnumid"
)

//获取临时票据根据公众号
func GetJsApiTicket(sessionId, pubnumid string) string {
	key := fmt.Sprintf(cache_jsapi_ticket, pubnumid)
	ticket, err := cache.Get(key)
	var b = true
	if err != nil {
		log.Error(err)
		b = false
	}
	if ticket == "" {
		b = false
	}
	if !b {
		pubnum := api.GetCacheTokenByPubnumId(pubnumid)
		wurl := fmt.Sprintf(jsapi_ticket, pubnum.Token)
		cjsn := webservice.SendGetJson(sessionId, wurl, "")
		if cjsn.Get("code").String() == "0" {
			ticket = cjsn.Get("ticket").String()
			cache.Set(key, ticket)
			cache.Expire(key, 7200)
			return ticket
		}
		return ""
	} else {
		return ticket
	}
}

//获得jssdk签名数据
func GenSign(sessionId, pubnumid, url string) map[string]string {
	//获得jsapi ticket
	ticket := GetJsApiTicket(sessionId, pubnumid)
	pubnum := api.GetCacheTokenByPubnumId(pubnumid)

	noncestr := uuid.NewV1().String()
	timestamp := date.NowUnixStr()
	appid := pubnum.Appid

	var b bytes.Buffer
	b.WriteString("jsapi_ticket=")
	b.WriteString(ticket)
	b.WriteString("&noncestr=")
	b.WriteString(noncestr)
	b.WriteString("&timestamp=")
	b.WriteString(timestamp)
	b.WriteString("&url=")
	b.WriteString(url)
	s1 := sha1.New()
	io.WriteString(s1, b.String())
	sign := fmt.Sprintf("%x", s1.Sum(nil))
	mp := make(map[string]string)
	mp["appid"] = appid
	mp["timestamp"] = timestamp
	mp["noncestr"] = noncestr
	mp["sign"] = sign
	log.Info(mp)
	return mp
}
