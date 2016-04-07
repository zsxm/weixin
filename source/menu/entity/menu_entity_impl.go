//scgen
package entity

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/zsxm/scgo/data"
	"github.com/zsxm/scgo/tools"
)

//----------------------MenuBean begin--------------------------------------

type MenuBean struct {
	bean  *Menu
	beans *Menus
}

func NewMenuBean() *MenuBean {
	e := &MenuBean{}
	return e
}

func (this *MenuBean) NewEntity() data.EntityInterface {
	return NewMenu()
}

func (this *MenuBean) NewEntitys(cap int) data.EntitysInterface {
	return NewMenus(cap)
}

func (this *MenuBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *MenuBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *MenuBean) Datas() *Menus {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *MenuBean) Table() data.TableInformation {
	return menuTableInformation
}

func (this *MenuBean) FieldNames() data.FieldNames {
	return menuFieldNames
}

func (this *MenuBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Menu)
}

func (this *MenuBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Menus)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Menus struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewMenus(cap int) *Menus {
	e := &Menus{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Menus) SetPage(page *data.Page) {
	this.page = page
}

func (this *Menus) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Menu))
}

func (this *Menus) Values() []data.EntityInterface {
	return this.datas
}

func (this *Menus) Len() int {
	return len(this.datas)
}

func (this *Menus) Get(i int) *Menu {
	return this.datas[i].(*Menu)
}

func (this *Menus) Table() data.TableInformation {
	return menuTableInformation
}

func (this *Menus) FieldNames() data.FieldNames {
	return menuFieldNames
}

func (this *Menus) JSON() string {
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

//----------------------MenuBean end--------------------------------------

//----------------------Menu begin--------------------------------------
func NewMenu() *Menu {
	return &Menu{}
}

func (this *Menu) Id() *data.String {
	return &this.id
}

func (this *Menu) SetId(value string) {
	this.id.SetValue(value)
}

func (this *Menu) Created() *data.Integer {
	return &this.created
}

func (this *Menu) SetCreated(value int) {
	this.created.SetValue(strconv.Itoa(value))
}

func (this *Menu) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Menu) Table() data.TableInformation {
	return menuTableInformation
}

func (this *Menu) FieldNames() data.FieldNames {
	return menuFieldNames
}

func (this *Menu) Field(filedName string) data.EntityField {
	switch tools.Lower(filedName) {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "created":
		return this.created.StructType()
	}
	return nil
}

func (this *Menu) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"created":%q`, this.created.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewMenu end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	menuTableInformation.SetTableName("menu")
	menuColumnsArr := []string{
		"id", "created", 
	}
	menuTableInformation.SetColumns(menuColumnsArr)
	menuFieldNamesArr := []string{
		"id", "created", 
	}
	menuFieldNames.SetNames(menuFieldNamesArr)
}

var menuTableInformation data.TableInformation
var menuFieldNames data.FieldNames

//----------------------init() end--------------------------------------
