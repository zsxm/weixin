package webservice_test

import (
	"fmt"
	"io"
	"testing"
	"weixin/source/util/webservice"
)

func TestWebSer(t *testing.T) {
	upload := "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s"
	mtype, token := "video", "vePyAjvSCkyBrBiLAVQnjdI8n0CLMeYcZVHaRakSfMiFl-62MoKo57WPRG95Thmny3v9l7osvjs171jUBrihB9406CoQFAr8cDC36aql0GoITAeAIAYJQ"
	fmt.Println("token", token)
	fmt.Println("type", mtype)
	wurl := fmt.Sprintf(upload, token, mtype)
	fmt.Println("url", wurl)
	body := make(map[string][]string)
	//body["media"] = []string{"file", "D:\\DOWN\\golangicon.jpg"}
	body["media"] = []string{"file", "D:\\DOWN\\TestMp4.mp4"}
	body["description"] = []string{"field", `{"title":"测试标题", "introduction":"视频素材的描述"}`}

	response, err := webservice.PostFile1(body, wurl)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	bye := make([]byte, response.ContentLength)
	c, err := response.Body.Read(bye)

	if err != nil && err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println("\nresponse body:", string(bye), " \nc :", c)
}
