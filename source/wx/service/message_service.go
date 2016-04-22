package service

import (
	pubnumApi "weixin/source/pubnum/api"
	"weixin/source/util"
	"weixin/source/util/consts"
	"weixin/source/websocket"
	"weixin/source/wx/entity"
	"weixin/source/wx/log"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/tools/date"
)

//验证token
func TokenChk(c chttp.Context) bool {
	var signature = c.Param("signature")
	var timestamp = c.Param("timestamp")
	var nonce = c.Param("nonce")
	var echostr = c.Param("echostr")
	//9066f01aff253ab20379b97a0e337d7d4b22910e 1461141818 1788485870 4496037518575137658
	log.Println(signature, timestamp, nonce, echostr)
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

//接收图片消息
//FromUserName 发送者
//ToUserName 接收者
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
	if r > 0 { //发送消息成功
		SaveMessage(string(imgMsg), reqMsg.MsgType, reqMsg.ToUserName, reqMsg.FromUserName)
	}
}

//接收文本消息
//FromUserName 发送者
//ToUserName 接收者
//content 回复内容
func Text(content string, reqMsg *entity.RequestMsg, c chttp.Context) {
	text := entity.ResponseMsgText{}
	text.CreateTime = date.NowUnixDuration()
	text.ToUserName = util.Value2CDATA(reqMsg.FromUserName)
	text.FromUserName = util.Value2CDATA(reqMsg.ToUserName)
	text.MsgType = util.Value2CDATA(consts.MSG_TEXT)
	//回复内容
	text.Content = util.Value2CDATA(content)

	userid := pubnumApi.GetPubNumUserid(reqMsg.ToUserName)
	msg := websocket.Message{Value: []byte(content), ReceId: userid}
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
	if r > 0 { //发送消息成功
		SaveMessage(string(textMsg), consts.MSG_TEXT, reqMsg.ToUserName, reqMsg.FromUserName)
	}
}

//保存消息,发送的和回复的
func SaveMessage(msg, mtype, formUser, toUser string) {
	e := entity.NewMessage()
	e.SetContent(msg)
	e.SetMtype(mtype)
	e.SetFormuser(formUser)
	e.SetTouser(toUser)
	e.Created().SetValue(date.NowUnixStr())
	sdbr, err := MessageService.Save(e)
	if err != nil {
		log.Error(err)
	}
	row, err := sdbr.RowsAffected()
	if err != nil {
		log.Error("SaveMessage: ", err)
	} else if row > 0 {
		log.Info("SaveMessage: ", "保存消息成功!")
	}
}
