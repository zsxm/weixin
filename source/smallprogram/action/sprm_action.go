//scgen
package action

import (
	"weixin/source/util"
	"weixin/source/wx/log"
	"weixin/source/wx/service"

	"github.com/zsxm/scgo/chttp"
)

//AppID(小程序ID)=wx834f2e0810975f86
//AppSecret(小程序密钥)=9090f1eb834d054356b65010e9caaefb
//Token(令牌)=a07fbd22a3bd11e5ada43010b3a0f
//EncodingAESKey=iVtJJ5WRj0hntnDU2k05qr0tItUUnJItaPMSgRByrtv
func init() {
	control.Add("/sprm/demo", sprm).Get()
}

func sprm(c chttp.Context) {
	b := service.TokenChk(c)
	if b && c.Method() == "POST" {
		msg, err := c.ReadBody()
		if err != nil {
			log.Error(err)
		}
		reqMsg, err := util.RequestMsg(msg)
		if err != nil {
			log.Error(err)
		}
		/*
			在发送方看待：
			FromUserName 发送者
			ToUserName 接收者
			反之倒过来
		*/
		log.Println("request msg type:", reqMsg.MsgType)
		log.Printf("%+v msg struct:\n", reqMsg)
	}
}
