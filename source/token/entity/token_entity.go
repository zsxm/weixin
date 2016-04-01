package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=weixin -moduleName=token -goSource=source
//go:@Table value=token comment=Token title=Token管理
type Token struct {

	//go:@Column value=id
	//go:@Identif
	id data.String

	//go:@Column value=token
	token data.String

	//go:@Column value=pubnumid
	pubnumid data.String

	//go:@Column value=created
	created data.Integer
}
