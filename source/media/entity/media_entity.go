package entity

import (
	"github.com/zsxm/scgo/data"
)

//go:generate $GOPATH/src/github.com/zsxm/scgo/tools/scgen/scgen.exe -fileDir=$GOFILE -projectDir=weixin -moduleName=media -goSource=source
//go:@Table value=media comment=素材 title=素材管理
type Media struct {

	//go:@Column value=id
	//go:@Identif
	id data.String

	//go:@Column value=mediaId
	mediaId data.String

	//go:@Column value=localName
	localName data.String

	//go:@Column value=ctype
	ctype data.String

	//go:@Column value=created
	created data.Integer

	//go:@Column value=url
	url data.String

	//go:@Column value=saveType
	saveType data.Integer

	//go:@Column value=title
	title data.String

	//go:@Column value=introduction
	introduction data.String
}
