//scgen
package entity

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/zsxm/scgo/data"
	"github.com/zsxm/scgo/tools"
)

//----------------------MediaBean begin--------------------------------------

type MediaBean struct {
	bean  *Media
	beans *Medias
}

func NewMediaBean() *MediaBean {
	e := &MediaBean{}
	return e
}

func (this *MediaBean) NewEntity() data.EntityInterface {
	return NewMedia()
}

func (this *MediaBean) NewEntitys(cap int) data.EntitysInterface {
	return NewMedias(cap)
}

func (this *MediaBean) Entity() data.EntityInterface {
	if this.bean == nil {
		return nil
	}
	return this.bean
}

func (this *MediaBean) Entitys() data.EntitysInterface {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *MediaBean) Datas() *Medias {
	if this.beans == nil {
		return nil
	}
	return this.beans
}

func (this *MediaBean) Table() data.TableInformation {
	return mediaTableInformation
}

func (this *MediaBean) FieldNames() data.FieldNames {
	return mediaFieldNames
}

func (this *MediaBean) SetEntity(bean data.EntityInterface) {
	this.bean = bean.(*Media)
}

func (this *MediaBean) SetEntitys(beans data.EntitysInterface) {
	this.beans = beans.(*Medias)
}

//------------------------------------------------------------

//------------------------------------------------------------
type Medias struct {
	datas []data.EntityInterface
	page  *data.Page
}

func NewMedias(cap int) *Medias {
	e := &Medias{}
	e.datas = make([]data.EntityInterface, 0, cap)
	return e
}

func (this *Medias) SetPage(page *data.Page) {
	this.page = page
}

func (this *Medias) Add(e data.EntityInterface) {
	this.datas = append(this.datas, e.(*Media))
}

func (this *Medias) Values() []data.EntityInterface {
	return this.datas
}

func (this *Medias) Len() int {
	return len(this.datas)
}

func (this *Medias) Get(i int) *Media {
	return this.datas[i].(*Media)
}

func (this *Medias) Table() data.TableInformation {
	return mediaTableInformation
}

func (this *Medias) FieldNames() data.FieldNames {
	return mediaFieldNames
}

func (this *Medias) JSON() string {
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

//----------------------MediaBean end--------------------------------------

//----------------------Media begin--------------------------------------
func NewMedia() *Media {
	return &Media{}
}

func (this *Media) Id() *data.String {
	return &this.id
}

func (this *Media) SetId(value string) {
	this.id.SetValue(value)
}

func (this *Media) MediaId() *data.String {
	return &this.mediaId
}

func (this *Media) SetMediaId(value string) {
	this.mediaId.SetValue(value)
}

func (this *Media) LocalName() *data.String {
	return &this.localName
}

func (this *Media) SetLocalName(value string) {
	this.localName.SetValue(value)
}

func (this *Media) Ctype() *data.String {
	return &this.ctype
}

func (this *Media) SetCtype(value string) {
	this.ctype.SetValue(value)
}

func (this *Media) Created() *data.Integer {
	return &this.created
}

func (this *Media) SetCreated(value int) {
	this.created.SetValue(strconv.Itoa(value))
}

func (this *Media) Url() *data.String {
	return &this.url
}

func (this *Media) SetUrl(value string) {
	this.url.SetValue(value)
}

func (this *Media) SaveType() *data.Integer {
	return &this.saveType
}

func (this *Media) SetSaveType(value int) {
	this.saveType.SetValue(strconv.Itoa(value))
}

func (this *Media) Title() *data.String {
	return &this.title
}

func (this *Media) SetTitle(value string) {
	this.title.SetValue(value)
}

func (this *Media) Introduction() *data.String {
	return &this.introduction
}

func (this *Media) SetIntroduction(value string) {
	this.introduction.SetValue(value)
}

func (this *Media) SetValue(filedName, value string) {
	this.Field(filedName).SetValue(value)
}

func (this *Media) Table() data.TableInformation {
	return mediaTableInformation
}

func (this *Media) FieldNames() data.FieldNames {
	return mediaFieldNames
}

func (this *Media) Field(filedName string) data.EntityField {
	switch tools.Lower(filedName) {
	case "id":
		this.id.SetPrimaryKey(true)
		return this.id.StructType()
	case "mediaid":
		return this.mediaId.StructType()
	case "localname":
		return this.localName.StructType()
	case "ctype":
		return this.ctype.StructType()
	case "created":
		return this.created.StructType()
	case "url":
		return this.url.StructType()
	case "savetype":
		return this.saveType.StructType()
	case "title":
		return this.title.StructType()
	case "introduction":
		return this.introduction.StructType()
	}
	return nil
}

func (this *Media) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")
	b.WriteString(fmt.Sprintf(`"id":%q`, this.id.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"mediaId":%q`, this.mediaId.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"localName":%q`, this.localName.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"ctype":%q`, this.ctype.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"created":%q`, this.created.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"url":%q`, this.url.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"saveType":%q`, this.saveType.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"title":%q`, this.title.Value()))
	b.WriteString(",")
	b.WriteString(fmt.Sprintf(`"introduction":%q`, this.introduction.Value()))
	b.WriteString("}")
	return b.String()
}

//----------------------NewMedia end--------------------------------------

//----------------------init() end--------------------------------------
func init() {
	mediaTableInformation.SetTableName("media")
	mediaColumnsArr := []string{
		"id", "mediaId", "localName", "ctype", "created", "url", "saveType", "title", "introduction", 
	}
	mediaTableInformation.SetColumns(mediaColumnsArr)
	mediaFieldNamesArr := []string{
		"id", "mediaId", "localName", "ctype", "created", "url", "saveType", "title", "introduction", 
	}
	mediaFieldNames.SetNames(mediaFieldNamesArr)
}

var mediaTableInformation data.TableInformation
var mediaFieldNames data.FieldNames

//----------------------init() end--------------------------------------
