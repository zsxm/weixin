package service

import (
	"weixin/source/util"
	"weixin/source/util/consts"
	"weixin/source/wx/entity"
	"weixin/source/wx/log"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

func Event(reqMsg *entity.RequestMsg, c chttp.Context) {
	switch reqMsg.Event.Event {
	case consts.EVENT_SUBSCRIBE:
		break
	case consts.EVENT_UNSUBSCRIBE:
		break
	case consts.EVENT_SCAN:
		break
	case consts.EVENT_LOCATION:
		break
	case consts.EVENT_CLICK:
		Click(reqMsg, c)
		break
	case consts.EVENT_VIEW:
		break
	}
}

//接收文本消息
func Click(reqMsg *entity.RequestMsg, c chttp.Context) {
	text := entity.ResponseMsgText{}
	text.CreateTime = date.NowUnixDuration()
	text.ToUserName = util.Value2CDATA(reqMsg.FromUserName)
	text.FromUserName = util.Value2CDATA(reqMsg.ToUserName)
	text.MsgType = util.Value2CDATA(consts.MSG_TEXT)
	text.Content = util.Value2CDATA("Hello , ", reqMsg.FromUserName)

	textMsg, err := util.ResponseMsg(text)
	if err != nil {
		log.Error(err)
	}
	c.Response().Header().Set("Content-Type", "text/xml")
	r, err := c.Write(textMsg)
	if err != nil {
		log.Error(err)
	}
	log.Info("send msg len", r)
}
