package webservice

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	pubNumApi "weixin/source/pubnum/api"
	"weixin/source/util"

	"github.com/zsxm/scgo/chttplib"
	"github.com/zsxm/scgo/cjson"
	"github.com/zsxm/scgo/data/cache"
	"github.com/zsxm/scgo/session"
)

const (
	//获取token
	token_url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

//接口 获取token
func GetWeiXinToKen(appid, secret string) *cjson.JSON {
	if appid != "" && secret != "" {
		url := fmt.Sprintf(token_url, appid, secret)
		cjson := SendGetJson("", url, "")
		return cjson
	}
	return nil
}

//Get请求发送数据,响应string
func SendGet(url, data string) string {
	request := chttplib.Get(url)
	var res string
	var err error
	if data != "" {
		request.Body(data)
	}
	loger.Info("URL:", url)
	loger.Info("DATA:", data)
	res, err = request.String()
	loger.Info("SendGet Response:", res)

	//res = `{"menu":{"button":[{"type":"click","name":"菜单1","key":"V1001_TODAY_MUSIC","sub_button":[]},{"name":"菜单2","sub_button":[{"type":"view","name":"搜索","url":"http:\/\/www.soso.com\/","sub_button":[]},{"type":"view","name":"视频","url":"http:\/\/v.qq.com\/","sub_button":[]},{"type":"click","name":"赞一下我们","key":"V1001_GOOD","sub_button":[]}]},{"name":"菜单三","sub_button":[{"type":"click","name":"子菜单名称","key":"keyabce","sub_button":[]}]}]}}`
	//res = `{"errcode":40013,"errmsg":"invalid appid"}`
	if err != nil {
		loger.Error(err)
	} else {
		return res
	}
	return ""
}

//Post请求发送数据,响应string
func SendPost(url, data string) string {
	request := chttplib.Post(url)
	if data != "" {
		request.Body(data)
	}
	loger.Info("URL:", url)
	loger.Info("DATA:", data)
	res, err := request.String()
	loger.Info("SendPost Response:", res)

	if err != nil {
		loger.Error(err)
	} else {
		return res
	}
	return ""
}

//通过sessionid 获取token并缓存
func GetToKen(sessionId string) (string, error) {
	if sessionId != "" {
		pubNum, pubNumId := CachePubNumBySessionId(sessionId)
		cjsn := GetWeiXinToKen(pubNum.Appid, pubNum.Secret)
		if cjsn != nil && pubNumId != "" {
			//缓存公众号基本信息
			pubNum.Token = cjsn.Get("access_token").String()
			//设置缓存
			pubNumApi.SetCachePubNum(pubNumId, pubNum)
			return pubNum.Token, nil
		}
		return "", errors.New("获取token失败,请检查是否启用了一个公众号!")
	}
	return "", errors.New("请登录!")
}

//获得缓存的公众号
func CachePubNumBySessionId(sessionId string) (pubNumApi.CachePubnum, string) {
	sessionMap, err := cache.HGetMap(session.SessionPrefix + sessionId)
	if err != nil {
		loger.Error(err)
	} else {
		userId := sessionMap.Get("id")                 //获取用户id
		pubNumId := pubNumApi.GetCachePubNumId(userId) //获取启用的公众号id
		pubNum := pubNumApi.CachePubNum(pubNumId)      //获取缓存的公众号
		fmt.Printf("%+v\n", pubNum)
		return pubNum, pubNumId //返回 缓存的公众号
	}
	return pubNumApi.CachePubnum{}, ""
}

//Get请求发送数据,响应json
func SendGetJson(sessionId, url, data string) *cjson.JSON {
	res := SendGet(url, data)
	loger.Info("send url", url, "data", data)
	var cjsn *cjson.JSON
	if res != "" {
		cjsn = cjson.JsonToMap(res)
		var code, codemsg string
		if cjsn.Size() > 0 {
			code = cjsn.Get("errcode").String()
			if code == "41001" || code == "42001" || code == "40014" || code == "40001" { //判断是否为token错误
				tk, err := GetToKen(sessionId)
				if err == nil {
					reg := regexp.MustCompile(`access_token.*?&{0,1}`)
					url = reg.ReplaceAllString(url, `access_token=`+tk+`&`)
					return SendGetJson(sessionId, url, data) //重新发送请求
				}
				loger.Error("SendGetJson", err)
				cjsn.Set("code", "130")
				cjsn.Set("codemsg", err.Error())
			} else if code != "" {
				cjsn.Set("code", code)
				ec, err := strconv.Atoi(code)
				if err != nil {
					loger.Error("errcode", code, err)
				} else {
					codemsg = util.ErrorMsg(ec)
					cjsn.Set("codemsg", codemsg)
				}
			} else {
				cjsn.Set("code", "0")
				cjsn.Set("codemsg", "ok")
			}
		}
		return cjsn
	}
	return cjsn
}

//Post请求发送数据,响应json
func SendPostJson(sessionId, url, data string) *cjson.JSON {
	res := SendPost(url, data)
	loger.Info("send url", url, "data", data)
	var cjsn *cjson.JSON
	if res != "" {
		cjsn = cjson.JsonToMap(res)
		var code, codemsg string
		if cjsn.Size() > 0 {
			code = cjsn.Get("errcode").String()
			if code == "41001" || code == "42001" || code == "40014" || code == "40001" { //判断是否为token错误
				tk, err := GetToKen(sessionId)
				if err == nil {
					reg := regexp.MustCompile(`access_token.*?&`)
					url = reg.ReplaceAllString(url, `access_token=`+tk+`&`)
					return SendPostJson(sessionId, url, data) //重新发送请求
				}
				loger.Error("SendGetJson", err)
				cjsn.Set("code", "130")
				cjsn.Set("codemsg", err.Error())
			} else if code != "" {
				cjsn.Set("code", code)
				ec, err := strconv.Atoi(code)
				if err != nil {
					loger.Error("errcode", code, err)
				} else {
					codemsg = util.ErrorMsg(ec)
					cjsn.Set("codemsg", codemsg)
				}
			} else {
				cjsn.Set("code", "0")
				cjsn.Set("codemsg", "ok")
			}
		}
		return cjsn
	}
	return cjsn
}

//上传媒体素材 带表单元素和媒体文件
//formField:map key=formName value=[]string{"type[file&field]","value"}
//return json str
func SendMediaUpload(formField map[string][]string, target_url string) string {
	response, err := PostFormFile(formField, target_url)
	if err != nil {
		loger.Error("PostFormFile:", err)
		return ""
	}
	defer response.Body.Close()
	bye := make([]byte, response.ContentLength)
	c, err := response.Body.Read(bye)

	if err != nil && err != io.EOF {
		loger.Error("Response Body Read:", err)
		return ""
	}
	loger.Info("Response Body Read Size:", c)
	return string(bye)
}

//上传媒体素材 带表单元素和媒体文件
//formField:map key=formName value=[]string{"type[file&field]","value"}
//return json object
func SendMediaUploadJson(formField map[string][]string, target_url string) *cjson.JSON {
	res := SendMediaUpload(formField, target_url)
	var cjsn *cjson.JSON
	if res != "" {
		cjsn = cjson.JsonToMap(res)
		var code, codemsg string
		if cjsn.Size() > 0 {
			code = cjsn.Get("errcode").String()
			if code != "" {
				cjsn.Set("code", code)
				ec, err := strconv.Atoi(code)
				if err != nil {
					loger.Error("errcode", code, err)
				} else {
					codemsg = util.ErrorMsg(ec)
					cjsn.Set("codemsg", codemsg)
				}
			} else {
				cjsn.Set("code", "0")
				cjsn.Set("codemsg", "ok")
			}
		}
		return cjsn
	}
	return cjsn
}
