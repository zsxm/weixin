package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=weixin -moduleName=user -goSource=source
//go:@Table value=user
type User struct {

	//go:@Column value=id
	//go:@Identif
	id data.String

	//go:@Column value=name
	name data.String

	//go:@Column value=nickname
	nickname data.String

	//go:@Column value=account
	account data.String

	//go:@Column value=passwd
	passwd data.String

	//go:@Column value=created
	created data.Integer
}
