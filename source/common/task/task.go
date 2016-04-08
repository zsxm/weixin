package task

import (
	"weixin/source/token/api"

	"github.com/zsxm/scgo/tools/cron"
)

func init() {
	cron.Add("cacheToken", api.Start)
}
