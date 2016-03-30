package main

import (
	_ "weixin/init"

	"github.com/zsxm/scgo/chttp"
)

func main() {
	chttp.Conf = &chttp.Config{
		Static: chttp.Mapping{
			Dir:    "static",
			Prefix: "/static",
			Ext:    []string{"js", "css", "gif", "jpg", "png", "ico", "map", "woff2", "ttf", "woff", "svg", "eot"},
		},
		Html: chttp.Mapping{
			Dir:    "template",
			Prefix: "/",
			Ext:    []string{"html"},
		},
		Error404: chttp.ErrorPage{
			Url:     "/error404",
			Message: "404 not page",
		},
		Error500: chttp.ErrorPage{
			Url:     "/error500",
			Message: "server 500 error",
		},
		Template: chttp.Template{
			Dir: "template",
		},
		Debug: true,
		Port:  ":8080",
	}
	chttp.Run()
}
