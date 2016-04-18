package service

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"weixin/source/common/log"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/cjson"
)

const (
	SUCCESS               = "SUCCESS"
	MAX_SIZE              = "MAX_SIZE"
	PERMISSION_DENIED     = "PERMISSION_DENIED"
	FAILED_CREATE_FILE    = "FAILED_CREATE_FILE"
	IO_ERROR              = "IO_ERROR"
	NOT_MULTIPART_CONTENT = "NOT_MULTIPART_CONTENT"
	PARSE_REQUEST_ERROR   = "PARSE_REQUEST_ERROR"
	NOTFOUND_UPLOAD_DATA  = "NOTFOUND_UPLOAD_DATA"
	NOT_ALLOW_FILE_TYPE   = "NOT_ALLOW_FILE_TYPE"
	INVALID_ACTION        = "INVALID_ACTION"
	CONFIG_ERROR          = "CONFIG_ERROR"
	PREVENT_HOST          = "PREVENT_HOST"
	CONNECTION_ERROR      = "CONNECTION_ERROR"
	REMOTE_FAIL           = "REMOTE_FAIL"
	NOT_DIRECTORY         = "NOT_DIRECTORY"
	NOT_EXIST             = "NOT_EXIST"
	ILLEGAL               = "ILLEGAL"
)

var (
	config   string
	cjsn     *cjson.JSON
	isAction = make(map[string]int)
)

func init() {
	fi, err := os.Open("conf/config.json")
	if err != nil {
		log.Error(err)
	}
	defer fi.Close()
	bye, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Error(err)
	}
	reg, err := regexp.Compile(`/\*{1,2}[\s\S]*?\*/`)
	if err != nil {
		log.Error(err)
	}
	config = reg.ReplaceAllString(string(bye), "")
	reg, err = regexp.Compile(`/\r|\n|\s|\t/`)
	if err != nil {
		log.Error(err)
	}
	config = reg.ReplaceAllString(config, "")
	cjsn = cjson.JsonToMap(config)
	if cjsn != nil {
		isAction["config"] = 0
		isAction[cjsn.Get("imageActionName").String()] = 1        /* 上传图片配置项 */
		isAction[cjsn.Get("scrawlActionName").String()] = 2       /* 涂鸦图片上传配置项 */
		isAction[cjsn.Get("videoActionName").String()] = 3        /* 上传视频配置 */
		isAction[cjsn.Get("fileActionName").String()] = 4         /* 上传文件配置 */
		isAction[cjsn.Get("imageManagerActionName").String()] = 5 /* 列出指定目录下的图片 */
		isAction[cjsn.Get("fileManagerActionName").String()] = 6  /* 列出指定目录下的文件 */
		isAction[cjsn.Get("snapscreenActionName").String()] = 7   /* 截图工具上传 */
		isAction[cjsn.Get("catcherActionName").String()] = 8      /* 抓取远程图片配置 */
	}
}

//获得配置信息
func Config() string {
	return config
}

func IsAction(action string) int {
	return isAction[action]
}

//上传
func Upload(c chttp.Context, isAction int) (string, error) {
	result := make(map[string]string)
	var config string
	switch isAction {
	case 1:
		config = Image()
		break
	case 2:
		config = Scrawl()
		break
	case 3:
		config = Video()
		break
	case 4:
		config = File()
		break
	}
	var jsn = cjson.JsonToMap(config)
	var mfi = c.MultiFile()
	result["original"] = mfi.FileName[0]
	result["type"] = mfi.FileSuffix[0]
	size, _ := strconv.Atoi(c.Param("size"))
	id := c.Param("id")
	name := c.Param("name")
	cMaxSize := jsn.Get("maxSize").Integer()
	if size > cMaxSize {
		result["state"] = MAX_SIZE
		return map2json(result), errors.New("文件大小超过限制")
	}
	log.Println(size, id, name, cMaxSize)
	var path = jsn.Get("savePath").String()
	var root = jsn.Get("rootPath").String()
	mfi.DirName = path
	err := mfi.Upload(root)
	if err != nil {
		result["state"] = IO_ERROR
		return map2json(result), err
	}
	result["state"] = SUCCESS
	result["name"] = mfi.FileNameId[0]
	result["url"] = mfi.DirName + mfi.FileNameId[0]
	result["size"] = strconv.FormatInt(mfi.Size[0], 10)

	//{"original":"demo.jpg","name":"demo.jpg","url":"\/server\/ueditor\/upload\/image\/demo.jpg","size":"99697","type":".jpg","state":"SUCCESS"}
	return map2json(result), nil
}

//文件列表
func FileList() {

}

//map 转 json
func map2json(data map[string]string) string {
	r, err := cjson.MapToJson(data)
	if err != nil {
		log.Error(err)
		return ""
	}
	return string(r)
}

/* 上传图片配置项 */
func Image() string {
	res := make(map[string]string)
	res["isBase64"] = "false"
	res["maxSize"] = cjsn.Get("imageMaxSize").String()
	res["allowFiles"] = cjsn.Get("imageAllowFiles").String()
	res["fieldName"] = cjsn.Get("imageFieldName").String()
	res["savePath"] = cjsn.Get("imagePathFormat").String()
	res["rootPath"] = "upload"
	jsn, err := cjson.MapToJson(res)
	if err != nil {
		log.Error(err)
	}
	return string(jsn)
}

/* 涂鸦图片上传配置项 */
func Scrawl() string {
	res := make(map[string]string)
	res["isBase64"] = "true"
	res["filename"] = "scrawl"
	res["maxSize"] = cjsn.Get("scrawlMaxSize").String()
	res["fieldName"] = cjsn.Get("scrawlFieldName").String()
	res["savePath"] = cjsn.Get("scrawlPathFormat").String()
	res["rootPath"] = "upload"
	jsn, err := cjson.MapToJson(res)
	if err != nil {
		log.Error(err)
	}
	return string(jsn)
}

/* 截图工具上传 */
func Snapscreen() string {
	return ""
}

/* 抓取远程图片配置 */
func Catcher() string {
	res := make(map[string]string)
	res["filename"] = "remote"
	res["filter"] = cjsn.Get("catcherLocalDomain").String()
	res["maxSize"] = cjsn.Get("catcherMaxSize").String()
	res["allowFiles"] = cjsn.Get("catcherAllowFiles").String()
	res["fieldName"] = cjsn.Get("catcherFieldName").String() + "[]"
	res["savePath"] = cjsn.Get("catcherPathFormat").String()
	jsn, err := cjson.MapToJson(res)
	if err != nil {
		log.Error(err)
	}
	return string(jsn)
}

/* 上传视频配置 */
func Video() string {
	res := make(map[string]string)
	res["maxSize"] = cjsn.Get("videoMaxSize").String()
	res["allowFiles"] = cjsn.Get("videoAllowFiles").String()
	res["fieldName"] = cjsn.Get("videoFieldName").String()
	res["savePath"] = cjsn.Get("videoPathFormat").String()
	res["rootPath"] = "upload"
	jsn, err := cjson.MapToJson(res)
	if err != nil {
		log.Error(err)
	}
	return string(jsn)
}

/* 上传文件配置 */
func File() string {
	res := make(map[string]string)
	res["isBase64"] = "false"
	res["maxSize"] = cjsn.Get("fileMaxSize").String()
	res["allowFiles"] = cjsn.Get("fileAllowFiles").String()
	res["fieldName"] = cjsn.Get("fileFieldName").String()
	res["savePath"] = cjsn.Get("filePathFormat").String()
	res["rootPath"] = "upload"
	jsn, err := cjson.MapToJson(res)
	if err != nil {
		log.Error(err)
	}
	return string(jsn)
}

/* 列出指定目录下的图片 */
func ImageManager() string {
	res := make(map[string]string)
	res["allowFiles"] = cjsn.Get("imageManagerAllowFiles").String()
	res["dir"] = cjsn.Get("imageManagerListPath").String()
	res["count"] = cjsn.Get("imageManagerListSize").String()
	jsn, err := cjson.MapToJson(res)
	if err != nil {
		log.Error(err)
	}
	return string(jsn)
}

/* 列出指定目录下的文件 */
func FileManager() string {
	res := make(map[string]string)
	res["allowFiles"] = cjsn.Get("fileManagerAllowFiles").String()
	res["dir"] = cjsn.Get("fileManagerListPath").String()
	res["count"] = cjsn.Get("fileManagerListSize").String()
	jsn, err := cjson.MapToJson(res)
	if err != nil {
		log.Error(err)
	}
	return string(jsn)
}
