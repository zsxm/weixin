package consts

const (
	//消息类型
	MSG_TEXT       = "text"       //文本
	MSG_IMAGE      = "image"      //图片
	MSG_VOICE      = "voice"      //语音
	MSG_VIDEO      = "video"      //视频
	MSG_SHORTVIDEO = "shortvideo" //小视频
	MSG_LOCATION   = "location"   //地理
	MSG_LINK       = "link"       //连接
	MSG_NEWS       = "news"       //图文
	MSG_MUSIC      = "music"      //音乐
	MSG_EVENT      = "event"      //事件

	//事件类型
	EVENT_SUBSCRIBE   = "subscribe"   //关注
	EVENT_UNSUBSCRIBE = "unsubscribe" //取消关注
	EVENT_SCAN        = "SCAN"        //已关注
	EVENT_LOCATION    = "LOCATION"    //上报地理位置
	EVENT_CLICK       = "CLICK"       //自定义菜单事件
	EVENT_VIEW        = "VIEW"        //点击菜单跳转链接时的事件
)
