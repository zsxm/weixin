package main

import (
	_ "weixin/init"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/config"
	"github.com/zsxm/scgo/data/cache"
	"github.com/zsxm/scgo/session"
)

func main() {
	//web项目配置
	config.Conf = &config.Config{
		TLSCrt: "E:/gopath/src/weixin/cert/test_server.crt",
		TLSKey: "E:/gopath/src/weixin/cert/test_server.key",
		Static: config.Mapping{
			Dir:    "static",
			Prefix: "/static",
			Ext:    []string{"js", "css", "gif", "jpg", "png", "ico", "map", "woff2", "ttf", "woff", "svg", "eot", "swf", "html"},
		},
		Html: config.Mapping{
			Dir:    "template",
			Prefix: "/",
			Ext:    []string{"html"},
		},
		Error404: config.ErrorPage{
			Url:     "/error404",
			Message: "404 not page",
		},
		Error500: config.ErrorPage{
			Url:     "/error500",
			Message: "server 500 error",
		},
		Template: config.Template{
			Dir: "template",
			Include: config.Include{
				Dir: "comm",
				Files: []string{
					"temp.header",
					"temp.left.menu",
					"temp.top",
				},
			},
		},
		Debug: true,
		//HttpPort: ":9093",
		Welcome: "/",
	}
	//redis 配置
	//redis 初始化化
	cache.Init(&cache.Config{
		Address:  "127.0.0.1:6379",
		Password: "foobared",
	})

	//session选项配置
	session.OptionsConfig = &session.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   30 * 60,
		Secure:   false,
		HttpOnly: true,
	}

	chttp.Run()
}
