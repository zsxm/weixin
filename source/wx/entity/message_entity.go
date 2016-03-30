package entity

import (
	"encoding/xml"
	"time"
)

//接收请求结构
type RequestMsg struct {
	Event
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`
	Content      string        `xml:"Content"`
	MediaId      string        `xml:"MediaId"`
	MsgId        int64         `xml:"MsgId"`
	Format       string        `xml:"Format"`
	Recognition  string        `xml:"Recognition"`
}

//事件
type Event struct {
	Event     string `xml:"Event"`
	EventKey  string `xml:"EventKey"`
	Ticket    string `xml:"Ticket"`
	Latitude  string `xml:"Latitude"`
	Longitude string `xml:"Longitude"`
	Precision string `xml:"Precision"`
}

//数据类型CDATA
type CDATAText struct {
	Text string `xml:",innerxml"`
}

//基础响应结构
type ResponseMsg struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   CDATAText     `xml:"ToUserName"`
	FromUserName CDATAText     `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      CDATAText     `xml:"MsgType"`
}

//文本响应
type ResponseMsgText struct {
	ResponseMsg
	Content CDATAText `xml:"Content"`
}

//图片响应
type ResponseImageMsg struct {
	ResponseMsg
	Image Image `xml:"Image"`
}

//图片
type Image struct {
	MediaId CDATAText `xml:"MediaId"`
}

//语音响应
type ResponseMsgVoice struct {
	ResponseMsg
	Voice Voice `xml:"Voice"`
}

//语音
type Voice struct {
	MediaId CDATAText `xml:"MediaId"`
}

//视频响应
type ResponseMsgVideo struct {
	ResponseMsg
	Video Video `xml:"Video"`
}

//视频
type Video struct {
	MediaId     CDATAText `xml:"MediaId"`
	Title       CDATAText `xml:"Title"`
	Description CDATAText `xml:"Description"`
}

//音乐响应
type ResponseMsgMusic struct {
	ResponseMsg
	Music Music `xml:"Music"`
}

//音乐
type Music struct {
	Title        CDATAText `xml:"Title"`
	Description  CDATAText `xml:"Description"`
	MusicUrl     CDATAText `xml:"MusicUrl"`
	HQMusicUrl   CDATAText `xml:"HQMusicUrl"`
	ThumbMediaId CDATAText `xml:"ThumbMediaId"`
}

//图文响应
type ResponseMsgArticles struct {
	ResponseMsg
	ArticleCount int      `xml:"ArticleCount"`
	Articles     Articles `xml:"Articles"`
}

//图文
type Articles struct {
	Item []Item `xml:"item"`
}

//图文集合
type Item struct {
	Title       CDATAText `xml:"Title"`
	Description CDATAText `xml:"Description"`
	PicUrl      CDATAText `xml:"PicUrl"`
	Url         CDATAText `xml:"Url"`
}
