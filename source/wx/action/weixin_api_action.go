package action

import (
	"weixin/source/util"
	"weixin/source/util/consts"
	"weixin/source/wx/api"
	"weixin/source/wx/log"
	"weixin/source/wx/service"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/weixin/demo", weixin)
	control.Add("/weixin/share", share)
}

//jssdk
func share(c chttp.Context) {
	pubnumid := c.Param("pubnumid")
	url := c.Param("url")
	mp := api.GenSign(pubnumid, url)
	c.JSON(mp, false)
}

func weixin(c chttp.Context) {
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
		switch reqMsg.MsgType {
		case consts.MSG_TEXT: //文本
			service.Text(reqMsg, c)
			break
		case consts.MSG_IMAGE: //图片
			service.Image(reqMsg, c)
			break
		case consts.MSG_VOICE: //语音
			break
		case consts.MSG_VIDEO: //视频
			break
		case consts.MSG_SHORTVIDEO: //小视频
			break
		case consts.MSG_LOCATION: //地理
			break
		case consts.MSG_LINK: //连接
			break
		case consts.MSG_NEWS: //图文
			break
		case consts.MSG_MUSIC: //音乐
			break
		case consts.MSG_EVENT: //事件
			service.Event(reqMsg, c)
			break
		}
	}
}
