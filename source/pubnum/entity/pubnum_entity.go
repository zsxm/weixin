package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=weixin -moduleName=pubnum -goSource=source
//go:@Table value=pubnum
type Pubnum struct {

	//go:@Column value=id
	//go:@Identif
	id data.String

	//go:@Column value=name
	name data.String

	//go:@Column value=appid
	appid data.String

	//go:@Column value=appsecret
	appsecret data.String

	//go:@Column value=token
	token data.String

	//go:@Column value=userid
	userid data.String

	//go:@Column value=created
	created data.Integer
}
