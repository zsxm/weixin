package service

import (
	"fmt"
	"weixin/source/util"
	"weixin/source/util/consts"
	"weixin/source/websocket"
	"weixin/source/wx/entity"
	"weixin/source/wx/log"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

//接收图片消息
func Image(reqMsg *entity.RequestMsg, c chttp.Context) {
	img := entity.ResponseImageMsg{}
	img.CreateTime = date.NowUnixDuration()
	img.ToUserName = util.Value2CDATA(reqMsg.FromUserName)
	img.FromUserName = util.Value2CDATA(reqMsg.ToUserName)
	img.MsgType = util.Value2CDATA(consts.MSG_IMAGE)
	img.Image = entity.Image{
		MediaId: util.Value2CDATA(reqMsg.MediaId),
	}

	imgMsg, err := util.ResponseMsg(img)
	if err != nil {
		log.Error(err)
	}
	c.Response().Header().Set("Content-Type", "text/xml")
	r, err := c.Write(imgMsg)
	if err != nil {
		log.Error(err)
	}
	log.Info("send msg len", r)
}

//接收文本消息
func Text(reqMsg *entity.RequestMsg, c chttp.Context) {
	text := entity.ResponseMsgText{}
	text.CreateTime = date.NowUnixDuration()
	text.ToUserName = util.Value2CDATA(reqMsg.FromUserName)
	text.FromUserName = util.Value2CDATA(reqMsg.ToUserName)
	text.MsgType = util.Value2CDATA(consts.MSG_TEXT)
	content := fmt.Sprint("Hello , ", reqMsg.ToUserName, " , ", reqMsg.Content)
	text.Content = util.Value2CDATA(content)

	msg := websocket.Message{Value: []byte(content), Key: reqMsg.ToUserName}
	websocket.H.Message <- msg

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

//验证token
func TokenChk(c chttp.Context) bool {
	var signature = c.Param("signature")
	var timestamp = c.Param("timestamp")
	var nonce = c.Param("nonce")
	var echostr = c.Param("echostr")
	b := util.SignValidation(signature, timestamp, nonce)
	if b {
		if echostr != "" {
			c.Write([]byte(echostr))
		}
	} else {
		c.Write([]byte("nil"))
	}
	return b
}
