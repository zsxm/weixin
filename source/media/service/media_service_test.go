package service_test

import (
	"testing"
	"weixin/source/media/service"

	"github.com/zsxm/scgo/data/cache"
)

func TestSync(t *testing.T) {
	//redis 配置
	//redis 初始化化
	cache.Init(&cache.Config{
		Address:  "127.0.0.1:6379",
		Password: "foobared",
	})
	service.GetMediaList("423F7D111E5A")
}
