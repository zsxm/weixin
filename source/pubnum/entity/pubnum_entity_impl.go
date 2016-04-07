//scgen
package entity

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/zsxm/scgo/data"
	"github.com/zsxm/scgo/tools"
)

//----------------------PubnumBean begin--------------------------------------

type PubnumBean struct {
	bean  *Pubnum
	beans *Pubnums
}

func NewPubnumBean() *PubnumBean {
	e := &PubnumBean{}
	return e
}

func (this *PubnumBean) NewEntity() data.EntityInterface {
	return NewPubnum()
}

func (this *PubnumBean) NewEntitys(cap int) data.EntitysInterface {
	return NewPubnums(cap)
}

func (this *PubnumBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *PubnumBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *PubnumBean) Datas() *Pubnums {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *PubnumBean) Table() data.TableInformation {
	return pubnumTableInformation
}

func (this *PubnumBean) FieldNames() data.FieldNames {
	return pubnumFieldNames
}

func (this *PubnumBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Pubnum)
}

func (this *PubnumBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Pubnums)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Pubnums struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewPubnums(cap int) *Pubnums {
	e := &Pubnums{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Pubnums) SetPage(page *data.Page) {
	this.page = page
}

func (this *Pubnums) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Pubnum))
}

func (this *Pubnums) Values() []data.EntityInterface {
	return this.datas
}

func (this *Pubnums) Len() int {
	return len(this.datas)
}

func (this *Pubnums) Get(i int) *Pubnum {
	return this.datas[i].(*Pubnum)
}

func (this *Pubnums) Table() data.TableInformation {
	return pubnumTableInformation
}

func (this *Pubnums) FieldNames() data.FieldNames {
	return pubnumFieldNames
}

func (this *Pubnums) JSON() string {
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

//----------------------PubnumBean end--------------------------------------

//----------------------Pubnum begin--------------------------------------
func NewPubnum() *Pubnum {
	return &Pubnum{}
}

func (this *Pubnum) Id() *data.String {
	return &this.id
}

func (this *Pubnum) SetId(value string) {
	this.id.SetValue(value)
}

func (this *Pubnum) Name() *data.String {
	return &this.name
}

func (this *Pubnum) SetName(value string) {
	this.name.SetValue(value)
}

func (this *Pubnum) Appid() *data.String {
	return &this.appid
}

func (this *Pubnum) SetAppid(value string) {
	this.appid.SetValue(value)
}

func (this *Pubnum) Appsecret() *data.String {
	return &this.appsecret
}

func (this *Pubnum) SetAppsecret(value string) {
	this.appsecret.SetValue(value)
}

func (this *Pubnum) Token() *data.String {
	return &this.token
}

func (this *Pubnum) SetToken(value string) {
	this.token.SetValue(value)
}

func (this *Pubnum) Userid() *data.String {
	return &this.userid
}

func (this *Pubnum) SetUserid(value string) {
	this.userid.SetValue(value)
}

func (this *Pubnum) Created() *data.Integer {
	return &this.created
}

func (this *Pubnum) SetCreated(value int) {
	this.created.SetValue(strconv.Itoa(value))
}

func (this *Pubnum) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Pubnum) Table() data.TableInformation {
	return pubnumTableInformation
}

func (this *Pubnum) FieldNames() data.FieldNames {
	return pubnumFieldNames
}

func (this *Pubnum) Field(filedName string) data.EntityField {
	switch tools.Lower(filedName) {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "name":
		return this.name.StructType()
	case "appid":
		return this.appid.StructType()
	case "appsecret":
		return this.appsecret.StructType()
	case "token":
		return this.token.StructType()
	case "userid":
		return this.userid.StructType()
	case "created":
		return this.created.StructType()
	}
	return nil
}

func (this *Pubnum) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"name":%q`, this.name.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"appid":%q`, this.appid.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"appsecret":%q`, this.appsecret.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"token":%q`, this.token.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"userid":%q`, this.userid.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"created":%q`, this.created.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewPubnum end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	pubnumTableInformation.SetTableName("pubnum")
	pubnumColumnsArr := []string{
		"id", "name", "appid", "appsecret", "token", "userid", "created", 
	}
	pubnumTableInformation.SetColumns(pubnumColumnsArr)
	pubnumFieldNamesArr := []string{
		"id", "name", "appid", "appsecret", "token", "userid", "created", 
	}
	pubnumFieldNames.SetNames(pubnumFieldNamesArr)
}

var pubnumTableInformation data.TableInformation
var pubnumFieldNames data.FieldNames

//----------------------init() end--------------------------------------
