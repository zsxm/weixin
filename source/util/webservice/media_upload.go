package webservice

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/zsxm/scgo/cjson"
)

//上传媒体文件
func PostFile(filename string, target_url string) (*http.Response, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	_, err := body_writer.CreateFormFile("media", filename)
	if err != nil {
		loger.Error("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
		loger.Error("error opening file")
		return nil, err
	}
	defer fh.Close()
	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()
	close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(close_string)
	// use multi-reader to defer the reading of the file data until writing to the socket buffer.
	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
		loger.Error("Error Stating file: ", filename)
		return nil, err
	}
	req, err := http.NewRequest("POST", target_url, request_reader)
	if err != nil {
		return nil, err
	}
	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	return http.DefaultClient.Do(req)
}

func GetFile(target_url string) *cjson.JSON {
	response, err := http.Get(target_url)
	if err != nil {
		loger.Error(err)
		return nil
	}
	defer response.Body.Close()
	ctype := response.Header["Content-Type"]
	var cjsn *cjson.JSON
	if response.StatusCode == 200 {
		if strings.Index(ctype[0], "text") != -1 {
			bye := make([]byte, response.ContentLength)
			_, err := response.Body.Read(bye)
			if err != nil {
				loger.Error(err)
				return nil
			}
			cjsn = cjson.JsonToMap(string(bye))
			cjsn.Set("code", "0")
			cjsn.Set("codemsg", "ok")
		} else {
			cjsn = &cjson.JSON{}
			cdisp := response.Header["Content-Disposition"]
			filename := cdisp[0]
			filename = filename[strings.Index(filename, "filename"):]
			filename = filename[10 : len(filename)-1]
			fi, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeType)
			if err != nil {
				loger.Error(err)
				return nil
			}
			defer fi.Close()
			_, err = io.Copy(fi, response.Body)
			cjsn.Set("code", "0")
			cjsn.Set("codemsg", "ok")
		}
	}
	return cjsn
}
