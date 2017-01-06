package action

import (
	"fmt"
	"net/url"
	"weixin/source/util"
	"weixin/source/util/consts"
	"weixin/source/util/webservice"
	"weixin/source/wx/api"
	"weixin/source/wx/log"
	"weixin/source/wx/service"

	"github.com/zsxm/scgo/cjson"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/weixin/demo", weixin)
	control.Add("/weixin/share", share)
	control.Add("/weixin/redirect", redirect).Get()
	control.Add("/weixin/redirect/test", redirectTest).Get()
}

//jssdk
func share(c chttp.Context) {
	pubnumid := c.Param("pubnumid")
	url := c.Param("url")
	mp := api.GenSign(c.Session().Id(), pubnumid, url)
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
		//保存发送来的消息
		switch reqMsg.MsgType {
		case consts.MSG_EVENT: //事件
			service.SaveMessage(string(msg), reqMsg.Event.Event, reqMsg.FromUserName, reqMsg.ToUserName)
			service.Event(msg, reqMsg, c)
			break
		default:
			service.SaveMessage(string(msg), reqMsg.MsgType, reqMsg.FromUserName, reqMsg.ToUserName)
			switch reqMsg.MsgType {
			case consts.MSG_TEXT: //文本
				service.Text("你好", reqMsg, c)
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
			}
		}
	}
}

//测试代码 begin
var (
	wxReadirect   = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s"
	appid         = "wx6c8a3beb5709e3d1"
	appsecret     = "d4624c36b6795d1d99dcf0547af5443d"
	scop_base     = "&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"
	scop_userinfo = "&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"
)

//测试页面跳转
//1 配置网页授权域名
//2 scope=snsapi_base和snsapi_userinfo区别
//	&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect
//	&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect
func redirect(c chttp.Context) {
	rtest := url.QueryEscape("http://1p48j81052.imwork.net/weixin/redirect/test") + scop_userinfo
	rurl := fmt.Sprintf(wxReadirect, appid, rtest)
	c.Redirect(rurl)
}

//redirect_uri
func redirectTest(c chttp.Context) {
	code := c.Param("code")
	//通过code换取网页授权access_token
	var wurl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	wurl = fmt.Sprintf(wurl, appid, appsecret, code)
	data := webservice.SendGet(wurl, "")
	cjsn := cjson.JsonToMap(data)

	access_token := cjsn.Get("access_token")
	refresh_token := cjsn.Get("refresh_token")
	openid := cjsn.Get("openid")
	scope := cjsn.Get("scope")
	unionid := cjsn.Get("unionid")
	log.Println("access_token", access_token)
	log.Println("refresh_token", refresh_token)
	log.Println("openid", openid)
	log.Println("scope", scope)
	log.Println("unionid", unionid)

	//检验授权凭证（access_token）是否有效
	wurl = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
	wurl = fmt.Sprintf(wurl, access_token, openid)
	data = webservice.SendGet(wurl, "")
	cjsn = cjson.JsonToMap(data)
	log.Println(cjsn.Data())

	//第四步：拉取用户信息(需scope为 snsapi_userinfo)
	wurl = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	wurl = fmt.Sprintf(wurl, access_token, openid)
	data = webservice.SendGet(wurl, "")
	cjsn = cjson.JsonToMap(data)
	log.Println(cjsn.Data())

	c.HTML("/weixin/weixin.readirect", cjsn.Data())
}

//测试代码 end
