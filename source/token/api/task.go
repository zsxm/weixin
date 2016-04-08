package api

import (
	pubNumApi "weixin/source/pubnum/api"
	"weixin/source/token/log"
	"weixin/source/token/service"

	"github.com/zsxm/scgo/tools/cron"
)

var (
	task_cron_spec = "0 0/30 * * * ?"
)

//获取每个公众号的token并缓存
func Start() {
	crn := cron.New()
	crn.AddFunc(task_cron_spec, func() {
		ToKenCacheTask()
	})
	crn.Start()
	log.Info("Pubnums List Token Task Start. OK")
}

//token 缓存任务
func ToKenCacheTask() {
	result, err := pubNumApi.PubNumIdList()
	if err != nil {
		log.Error(err)
		return
	}
	for _, v := range result.Data {
		pubnumid := v.Get("id")
		res := service.GetToKen(pubnumid, true)
		if res.Code != "0" {
			log.Error("ToKenCacheTask error:", res.Codemsg)
		}
	}
}
