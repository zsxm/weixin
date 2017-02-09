//scgen
package action

import (
	"strconv"
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
	control.Add("/sprm/list", list).Get()
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
		log.Println("request msg type:", reqMsg.MsgType)
		log.Printf("%+v msg struct:\n", reqMsg)
	}
	log.Println(c.Method())
}

func list(c chttp.Context) {
	datas := make([]interface{}, 0, 10)
	for i := 0; i < 10; i++ {
		data := make(map[string]string)
		data["id"] = strconv.Itoa(i + 1)
		data["name"] = "张" + strconv.Itoa(i+1)
		data["age"] = strconv.Itoa(i + i)
		datas = append(datas, data)
	}
	c.JSON(datas, false)
}
