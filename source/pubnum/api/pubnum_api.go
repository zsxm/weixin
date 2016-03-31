//scgen
package api

import (
	"weixin/source/pubnum/entity"
	"weixin/source/pubnum/service"
)

func GetPubnum(id string) *entity.Pubnum {
	pubnum := entity.NewPubnum()
	pubnum.Id().SetValue(id)
	pubnum.Id().FieldExp().Eq().And()
	service.PubnumService.SelectOne(pubnum)
	return pubnum
}
