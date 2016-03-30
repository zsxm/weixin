package util

import (
	"bytes"
	"weixin/source/wx/entity"

	"github.com/zsxm/scgo/cxml"
)

func RequestMsg(msg []byte) (*entity.RequestMsg, error) {
	recMsg := &entity.RequestMsg{}
	err := cxml.Unmarshal(recMsg, msg)
	if err != nil {
		return nil, err
	}
	return recMsg, nil
}

func ResponseMsg(e cxml.CxmlInterface) ([]byte, error) {
	b, err := cxml.Marshal(e)
	return b, err
}

func Value2CDATA(vals ...string) entity.CDATAText {
	sb := bytes.Buffer{}
	for _, v := range vals {
		sb.WriteString(v)
	}
	return entity.CDATAText{"<![CDATA[" + sb.String() + "]]>"}
}
