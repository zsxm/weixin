package util_test

import (
	"fmt"
	"testing"
	"weixin/source/common/entity"
	"weixin/source/common/util"
)

func TestUtil(t *testing.T) {
	text := entity.ResponseMsgText{}
	text.MsgType = util.Value2CDATA("text")
	text.CreateTime = 1231232312
	text.FromUserName = util.Value2CDATA("FromUserName")
	text.ToUserName = util.Value2CDATA("ToUserName")
	text.Content = util.Value2CDATA("内容")
	textMsg, err := util.ResponseMsg(text)
	fmt.Println(string(textMsg), err)
	//	img := entity.ResponseImageMsg{}
	//	img.MsgType = "text"
	//	img.CreateTime = 1231232312
	//	img.FromUserName = "FromUserName"
	//	img.ToUserName = "ToUserName"
	//	img.Image = entity.Image{
	//		MediaId: "123123123",
	//	}
	//	imaegMsg, err := util.ResponseMsg(img)
	//	fmt.Println(string(imaegMsg), err)

	//	art := entity.ResponseMsgArticles{}
	//	art.MsgType = "text"
	//	art.CreateTime = 1231232312
	//	art.FromUserName = "FromUserName"
	//	art.ToUserName = "ToUserName"

	//	art.ArticleCount = 2
	//	art.Articles = entity.Articles{
	//		Item: []entity.Item{
	//			entity.Item{
	//				Title:       "名称",
	//				Description: "描述",
	//				PicUrl:      "www.baidu.com/pic",
	//				Url:         "www.baidu.com",
	//			},
	//			entity.Item{
	//				Title:       "名称1",
	//				Description: "描述1",
	//				PicUrl:      "www.baidu.com/pic1",
	//				Url:         "www.baidu.com1",
	//			},
	//		},
	//	}
	//	artMsg, err := util.ResponseMsg(art)
	//	fmt.Println(string(artMsg), err)
}
