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
	if data != "" {
		request.Body(data)
	}
	res, err := request.String()

	//res = `{"access_token":"Y93-GXoRJAb33tUwiQIkZjm2MqmWlA3jQ28rGXPZV-0clPNM5WYBdImNHalz9JDq3lZKAlg27uBD5pMgLDMY4txxx6fCSJseHFi_bdvctiDcyvV4C3rFhqd7y9xnAKcdSCRcADADNC","expires_in":7200}`
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
	res, err := request.String()

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
