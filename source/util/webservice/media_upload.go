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

//上传带表单元素和媒体文件
//formField:map key=formName value=[]string{"type[file&field]","value"}
func PostFormFile(formField map[string][]string, target_url string) (*http.Response, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	for k, val := range formField {
		var err error
		if len(val) == 2 {
			var value = val[1]
			switch val[0] {
			case "field":
				err = body_writer.WriteField(k, value)
				break
			case "file":
				_, err = body_writer.CreateFormFile(k, value)
				fh, err := os.Open(value)
				if err != nil {
					loger.Error("error opening file")
					return nil, err
				}
				defer fh.Close()
				body_buf.ReadFrom(fh)
				break
			default:
				loger.Error(fmt.Errorf("form field name %s value type", k))
			}
			if err != nil {
				loger.Error("error writing to buffer")
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("form field name %s values lenght != 2", k)
		}
	}
	boundary := body_writer.Boundary()
	close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	body_buf.WriteString(close_string)
	request_reader := io.MultiReader(body_buf)

	req, err := http.NewRequest("POST", target_url, request_reader)
	if err != nil {
		fmt.Println("err")
		return nil, err
	}

	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = int64(body_buf.Len())

	return http.DefaultClient.Do(req)
}

//上传媒体文件
func PostFile(filename string, target_url string) (*http.Response, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	_, err := body_writer.CreateFormFile("media", filename)
	if err != nil {
		loger.Error("error writing to buffer")
		return nil, err
	}

	fh, err := os.Open(filename)
	if err != nil {
		loger.Error("error opening file")
		return nil, err
	}
	defer fh.Close()
	boundary := body_writer.Boundary()
	close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(close_string)
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
