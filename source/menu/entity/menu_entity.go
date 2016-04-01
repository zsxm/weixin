package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=weixin -moduleName=menu -goSource=source
//go:@Table value=menu
type Menu struct {

	//go:@Column value=id
	//go:@Identif
	id data.String

	//go:@Column value=created
	created data.Integer
}
