//scgen
package entity

import (
	"bytes"
	"fmt"
	"github.com/zsxm/scgo/data"
	"strconv"
)

//----------------------TokenBean begin--------------------------------------

type TokenBean struct {
	bean  *Token
	beans *Tokens
}

func NewTokenBean() *TokenBean {
	e := &TokenBean{}
	return e
}

func (this *TokenBean) NewEntity() data.EntityInterface {
	return NewToken()
}

func (this *TokenBean) NewEntitys(cap int) data.EntitysInterface {
	return NewTokens(cap)
}

func (this *TokenBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *TokenBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *TokenBean) Datas() *Tokens {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *TokenBean) Table() data.TableInformation {
	return tokenTableInformation
}

func (this *TokenBean) FieldNames() data.FieldNames {
	return tokenFieldNames
}

func (this *TokenBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Token)
}

func (this *TokenBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Tokens)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Tokens struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewTokens(cap int) *Tokens {
	e := &Tokens{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Tokens) SetPage(page *data.Page) {
	this.page = page
}

func (this *Tokens) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Token))
}

func (this *Tokens) Values() []data.EntityInterface {
	return this.datas
}

func (this *Tokens) Len() int {
	return len(this.datas)
}

func (this *Tokens) Get(i int) *Token {
	return this.datas[i].(*Token)
}

func (this *Tokens) Table() data.TableInformation {
	return tokenTableInformation
}

func (this *Tokens) FieldNames() data.FieldNames {
	return tokenFieldNames
}

func (this *Tokens) JSON() string {
	var wr bytes.Buffer
	wr.WriteString("[")
	for i, v := range this.Values() {
		if i > 0 {
			wr.WriteString(",")
		}
		wr.WriteString(v.JSON())
	}
	wr.WriteString("]")
	return wr.String()
}

//----------------------TokenBean end--------------------------------------

//----------------------Token begin--------------------------------------
func NewToken() *Token {
	return &Token{}
}

func (this *Token) Id() *data.String {
	return &this.id
}

func (this *Token) SetId(value string) {
	this.id.SetValue(value)
}

func (this *Token) Token() *data.String {
	return &this.token
}

func (this *Token) SetToken(value string) {
	this.token.SetValue(value)
}

func (this *Token) Pubnumid() *data.String {
	return &this.pubnumid
}

func (this *Token) SetPubnumid(value string) {
	this.pubnumid.SetValue(value)
}

func (this *Token) Created() *data.Integer {
	return &this.created
}

func (this *Token) SetCreated(value int) {
	this.created.SetValue(strconv.Itoa(value))
}

func (this *Token) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Token) Table() data.TableInformation {
	return tokenTableInformation
}

func (this *Token) FieldNames() data.FieldNames {
	return tokenFieldNames
}

func (this *Token) Field(filedName string) data.EntityField {
	switch filedName {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "token":
		return this.token.StructType()
	case "pubnumid":
		return this.pubnumid.StructType()
	case "created":
		return this.created.StructType()
	}
	return nil
}

func (this *Token) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"token":%q`, this.token.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"pubnumid":%q`, this.pubnumid.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"created":%q`, this.created.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewToken end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	tokenTableInformation.SetTableName("token")
	tokenColumnsArr := []string{
		"id", "token", "pubnumid", "created", 
	}
	tokenTableInformation.SetColumns(tokenColumnsArr)
	tokenFieldNamesArr := []string{
		"id", "token", "pubnumid", "created", 
	}
	tokenFieldNames.SetNames(tokenFieldNamesArr)
}

var tokenTableInformation data.TableInformation
var tokenFieldNames data.FieldNames

//----------------------init() end--------------------------------------
