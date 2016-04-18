package action

import (
	"weixin/source/common/log"
	"weixin/source/common/service"

	"github.com/zsxm/scgo/chttp"
)

func init() {
	control.Add("/comm/ueditor/ctrl", ueditorAction)
}

func ueditorAction(c chttp.Context) {
	action := c.Param("action")
	var isAction = service.IsAction(action)
	log.Info("action =", action, isAction)
	switch isAction {
	case 0:
		c.JSON(service.Config(), false)
		break
	case 1, 2, 3, 4:
		resJson, err := service.Upload(c, isAction)
		if err != nil {
			log.Error(err)
		}
		//var resJson = `{"original":"demo.jpg","name":"demo.jpg","url":"\/server\/ueditor\/upload\/image\/demo.jpg","size":"99697","type":".jpg","state":"SUCCESS"}`
		c.JSON(resJson, false)
		break
	case 5, 6:
		var resJson = `{"state":"SUCCESS","list":[{"url":"\/server\/ueditor\/upload\/image\/3.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/21.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/19.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/1.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/2.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/6.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/10.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/13.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/4 2.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/22.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/26.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/17.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/11.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/demo.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/16.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/9.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/4.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/20.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/12.jpg","mtime":1456378798},{"url":"\/server\/ueditor\/upload\/image\/18.jpg","mtime":1456378798}],"start":"0","total":29}`
		c.JSON(resJson, false)
		break
	}

}
