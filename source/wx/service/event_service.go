package service

import (
	"weixin/source/util/consts"
	"weixin/source/wx/entity"
	"weixin/source/wx/log"

	"github.com/zsxm/scgo/chttp"
)

func Event(msg []byte, reqMsg *entity.RequestMsg, c chttp.Context) {
	switch reqMsg.Event.Event {
	case consts.EVENT_SUBSCRIBE:
		SubScribe(reqMsg, c)
		break
	case consts.EVENT_UNSUBSCRIBE:
		UnSubScribe(reqMsg, c)
		break
	case consts.EVENT_SCAN:
		break
	case consts.EVENT_LOCATION:
		break
	case consts.EVENT_CLICK:
		break
	case consts.EVENT_VIEW:
		break
	}
}

//关注
func SubScribe(reqMsg *entity.RequestMsg, c chttp.Context) {
	//回复关注欢迎语
	log.Info("关注")
	Text("欢迎关注杨红岩微信公众号", reqMsg, c)
}

//取消关注
func UnSubScribe(reqMsg *entity.RequestMsg, c chttp.Context) {
	log.Info("有人取消关注")
}
