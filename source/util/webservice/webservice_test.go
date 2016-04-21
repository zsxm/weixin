package webservice_test

import (
	"fmt"
	"testing"
	"weixin/source/util/webservice"

	"github.com/zsxm/scgo/data/cache"
)

func TestWebService(t *testing.T) {
	webservice.GetToKen("F1549AB3E95449B18872D71E900B8344")
}

func init() {
	fmt.Println("a")
	cache.Init(&cache.Config{
		Address:  "127.0.0.1:6379",
		Password: "foobared",
	})
}
