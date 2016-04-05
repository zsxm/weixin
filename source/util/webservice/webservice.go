package webservice

import (
	"strconv"
	"weixin/source/util"

	"github.com/zsxm/scgo/chttplib"
	"github.com/zsxm/scgo/cjson"
)

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

//Get请求发送数据,响应json
func SendGetJson(url, data string) *cjson.JSON {
	res := SendGet(url, data)
	loger.Info("send url", url, "data", data)
	if res != "" {
		cjs := cjson.JsonToMap(res)
		var code, codemsg string
		if cjs.Size() > 0 {
			code = cjs.Get("errcode").String()
			if code != "" {
				cjs.Set("code", code)
				ec, err := strconv.Atoi(code)
				if err != nil {
					loger.Error("errcode", code, err)
				} else {
					codemsg = util.ErrorMsg(ec)
					cjs.Set("codemsg", codemsg)
				}
			} else {
				cjs.Set("code", "0")
				cjs.Set("codemsg", "ok")
			}
		}
		return cjs
	}
	return nil
}

//Post请求发送数据,响应json
func SendPostJson(url, data string) *cjson.JSON {
	res := SendPost(url, data)
	loger.Info("send url", url, "data", data)
	if res != "" {
		cjs := cjson.JsonToMap(res)
		var code, codemsg string
		if cjs.Size() > 0 {
			code = cjs.Get("errcode").String()
			if code != "" {
				cjs.Set("code", code)
				ec, err := strconv.Atoi(code)
				if err != nil {
					loger.Error("errcode", code, err)
				} else {
					codemsg = util.ErrorMsg(ec)
					cjs.Set("codemsg", codemsg)
				}
			} else {
				cjs.Set("code", "0")
				cjs.Set("codemsg", "ok")
			}
		}
		return cjs
	}
	return nil
}

func SendGetXml(url, xmlData string) {

}
