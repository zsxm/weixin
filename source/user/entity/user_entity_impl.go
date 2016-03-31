//scgen
package entity

import (
	"bytes"
	"fmt"
	"github.com/zsxm/scgo/data"
	"strconv"
)

//----------------------UserBean begin--------------------------------------

type UserBean struct {
	bean  *User
	beans *Users
}

func NewUserBean() *UserBean {
	e := &UserBean{}
	return e
}

func (this *UserBean) NewEntity() data.EntityInterface {
	return NewUser()
}

func (this *UserBean) NewEntitys(cap int) data.EntitysInterface {
	return NewUsers(cap)
}

func (this *UserBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *UserBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *UserBean) Datas() *Users {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *UserBean) Table() data.TableInformation {
	return userTableInformation
}

func (this *UserBean) FieldNames() data.FieldNames {
	return userFieldNames
}

func (this *UserBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*User)
}

func (this *UserBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Users)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Users struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewUsers(cap int) *Users {
	e := &Users{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Users) SetPage(page *data.Page) {
	this.page = page
}

func (this *Users) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*User))
}

func (this *Users) Values() []data.EntityInterface {
	return this.datas
}

func (this *Users) Len() int {
	return len(this.datas)
}

func (this *Users) Get(i int) *User {
	return this.datas[i].(*User)
}

func (this *Users) Table() data.TableInformation {
	return userTableInformation
}

func (this *Users) FieldNames() data.FieldNames {
	return userFieldNames
}

func (this *Users) JSON() string {
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

//----------------------UserBean end--------------------------------------

//----------------------User begin--------------------------------------
func NewUser() *User {
	return &User{}
}

func (this *User) Id() *data.String {
	return &this.id
}

func (this *User) SetId(value string) {
	this.id.SetValue(value)
}

func (this *User) Name() *data.String {
	return &this.name
}

func (this *User) SetName(value string) {
	this.name.SetValue(value)
}

func (this *User) Nickname() *data.String {
	return &this.nickname
}

func (this *User) SetNickname(value string) {
	this.nickname.SetValue(value)
}

func (this *User) Account() *data.String {
	return &this.account
}

func (this *User) SetAccount(value string) {
	this.account.SetValue(value)
}

func (this *User) Passwd() *data.String {
	return &this.passwd
}

func (this *User) SetPasswd(value string) {
	this.passwd.SetValue(value)
}

func (this *User) Created() *data.Integer {
	return &this.created
}

func (this *User) SetCreated(value int) {
	this.created.SetValue(strconv.Itoa(value))
}

func (this *User) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *User) Table() data.TableInformation {
	return userTableInformation
}

func (this *User) FieldNames() data.FieldNames {
	return userFieldNames
}

func (this *User) Field(filedName string) data.EntityField {
	switch filedName {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "name":
		return this.name.StructType()
	case "nickname":
		return this.nickname.StructType()
	case "account":
		return this.account.StructType()
	case "passwd":
		return this.passwd.StructType()
	case "created":
		return this.created.StructType()
	}
	return nil
}

func (this *User) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"name":%q`, this.name.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"nickname":%q`, this.nickname.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"account":%q`, this.account.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"passwd":%q`, this.passwd.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"created":%q`, this.created.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewUser end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	userTableInformation.SetTableName("user")
	userColumnsArr := []string{
		"id", "name", "nickname", "account", "passwd", "created", 
	}
	userTableInformation.SetColumns(userColumnsArr)
	userFieldNamesArr := []string{
		"id", "name", "nickname", "account", "passwd", "created", 
	}
	userFieldNames.SetNames(userFieldNamesArr)
}

var userTableInformation data.TableInformation
var userFieldNames data.FieldNames

//----------------------init() end--------------------------------------
