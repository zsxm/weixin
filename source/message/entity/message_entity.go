package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=weixin -moduleName=message -goSource=source
//go:@Table value=message comment=消息 title=消息管理
type Message struct {

	//go:@Column value=id
	//go:@Identif
	id data.String

	//go:@Column value=mtype
	mtype data.String

	//go:@Column value=content
	content data.String

	//go:@Column value=formuser
	formuser data.String

	//go:@Column value=touser
	touser data.String

	//go:@Column value=created
	created data.Integer
}
