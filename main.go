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
		Static: config.Mapping{
			Dir:    "static",
			Prefix: "/static",
			Ext:    []string{"js", "css", "gif", "jpg", "png", "ico", "map", "woff2", "ttf", "woff", "svg", "eot"},
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
					"header",
					"left.menu",
				},
			},
		},
		Debug:   true,
		Port:    ":8080",
		Welcome: "/",
	}
	//redis 配置
	cache.Conf = &cache.Config{
		Address:  "10.100.130.54:6379",
		Password: "foobared",
	}
	//redis 初始化化
	cache.Init(*cache.Conf)

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
