package service

import (
	"weixin/source/common/consts"
	"weixin/source/common/entity"

	"github.com/zsxm/scgo/chttp"
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
		break
	case consts.EVENT_VIEW:
		break
	}
}
