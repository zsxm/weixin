package webservice

import (
	"strconv"
	"weixin/source/util"

	"github.com/zsxm/scgo/chttplib"
	"github.com/zsxm/scgo/cjson"
)

func SendGet(url, jsonData string) string {
	request := chttplib.Get(url)
	if jsonData != "" {
		request.Body(jsonData)
	}
	res, err := request.String()

	res = `{"access_token":"Y93-GXoRJAb33tUwiQIkZjm2MqmWlA3jQ28rGXPZV-0clPNM5WYBdImNHalz9JDq3lZKAlg27uBD5pMgLDMY4txxx6fCSJseHFi_bdvctiDcyvV4C3rFhqd7y9xnAKcdSCRcADADNC","expires_in":7200}`
	//res = `{"errcode":40013,"errmsg":"invalid appid"}`
	if err != nil {
		loger.Error(err)
	} else {
		return res
	}
	return ""
}

func SendPost(url, jsonData string) string {
	request := chttplib.Post(url)
	if jsonData != "" {
		request.Body(jsonData)
	}
	res, err := request.String()

	if err != nil {
		loger.Error(err)
	} else {
		return res
	}
	return ""
}

func SendGetJson(url, jsonData string) *cjson.JSON {
	res := SendGet(url, jsonData)
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

func SendPostJson(url, jsonData string) *cjson.JSON {
	res := SendPost(url, jsonData)
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
