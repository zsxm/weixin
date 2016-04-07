//scgen
package service_test

import (
	"fmt"
	"testing"
	"weixin/source/util/webservice"
)

const (
	upload = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
)

func TestUpload(t *testing.T) {
	token, image := "Yefumfh9-Spq2Mu4tST3vftUCyrNJFSKKEPPrR_WQG-p_D-HyN9rGw4e_622b4SLwUZq4Nb_VgQvlIiaiEOoyQhURpb2gL5oSKDbKX90V7lyDpYerqT0dVcnhayT4JAvPTAaAFALPW", "image"
	wurl := fmt.Sprintf(upload, token, image)
	response, err := webservice.PostFile(`D:\DOWN\golangicon.jpg`, wurl)
	defer response.Body.Close()
	fmt.Println("response", err)
	be := make([]byte, 256)
	response.Body.Read(be)
	fmt.Println(string(be))
}
