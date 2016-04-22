//scgen
package entity

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/zsxm/scgo/data"
	"github.com/zsxm/scgo/tools"
)

//----------------------MessageBean begin--------------------------------------

type MessageBean struct {
	bean  *Message
	beans *Messages
}

func NewMessageBean() *MessageBean {
	e := &MessageBean{}
	return e
}

func (this *MessageBean) NewEntity() data.EntityInterface {
	return NewMessage()
}

func (this *MessageBean) NewEntitys(cap int) data.EntitysInterface {
	return NewMessages(cap)
}

func (this *MessageBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *MessageBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *MessageBean) Datas() *Messages {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *MessageBean) Table() data.TableInformation {
	return messageTableInformation
}

func (this *MessageBean) FieldNames() data.FieldNames {
	return messageFieldNames
}

func (this *MessageBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Message)
}

func (this *MessageBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Messages)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Messages struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewMessages(cap int) *Messages {
	e := &Messages{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Messages) SetPage(page *data.Page) {
	this.page = page
}

func (this *Messages) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Message))
}

func (this *Messages) Values() []data.EntityInterface {
	return this.datas
}

func (this *Messages) Len() int {
	return len(this.datas)
}

func (this *Messages) Get(i int) *Message {
	return this.datas[i].(*Message)
}

func (this *Messages) Table() data.TableInformation {
	return messageTableInformation
}

func (this *Messages) FieldNames() data.FieldNames {
	return messageFieldNames
}

func (this *Messages) JSON() string {
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

//----------------------MessageBean end--------------------------------------

//----------------------Message begin--------------------------------------
func NewMessage() *Message {
	return &Message{}
}

func (this *Message) Id() *data.String {
	return &this.id
}

func (this *Message) SetId(value string) {
	this.id.SetValue(value)
}

func (this *Message) Mtype() *data.String {
	return &this.mtype
}

func (this *Message) SetMtype(value string) {
	this.mtype.SetValue(value)
}

func (this *Message) Content() *data.String {
	return &this.content
}

func (this *Message) SetContent(value string) {
	this.content.SetValue(value)
}

func (this *Message) Formuser() *data.String {
	return &this.formuser
}

func (this *Message) SetFormuser(value string) {
	this.formuser.SetValue(value)
}

func (this *Message) Touser() *data.String {
	return &this.touser
}

func (this *Message) SetTouser(value string) {
	this.touser.SetValue(value)
}

func (this *Message) Created() *data.Integer {
	return &this.created
}

func (this *Message) SetCreated(value int) {
	this.created.SetValue(strconv.Itoa(value))
}

func (this *Message) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Message) Table() data.TableInformation {
	return messageTableInformation
}

func (this *Message) FieldNames() data.FieldNames {
	return messageFieldNames
}

func (this *Message) Field(filedName string) data.EntityField {
	switch tools.Lower(filedName) {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "mtype":
		return this.mtype.StructType()
	case "content":
		return this.content.StructType()
	case "formuser":
		return this.formuser.StructType()
	case "touser":
		return this.touser.StructType()
	case "created":
		return this.created.StructType()
	}
	return nil
}

func (this *Message) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"mtype":%q`, this.mtype.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"content":%q`, this.content.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"formuser":%q`, this.formuser.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"touser":%q`, this.touser.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"created":%q`, this.created.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewMessage end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	messageTableInformation.SetTableName("message")
	messageColumnsArr := []string{
		"id", "mtype", "content", "formuser", "touser", "created", 
	}
	messageTableInformation.SetColumns(messageColumnsArr)
	messageFieldNamesArr := []string{
		"id", "mtype", "content", "formuser", "touser", "created", 
	}
	messageFieldNames.SetNames(messageFieldNamesArr)
}

var messageTableInformation data.TableInformation
var messageFieldNames data.FieldNames

//----------------------init() end--------------------------------------
