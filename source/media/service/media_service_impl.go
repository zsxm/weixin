//scgen
package service

import (
	"database/sql"
	"github.com/zsxm/scgo/data"
	"github.com/zsxm/scgo/data/scdb"
)

var MediaService = mediaService{
	repository: scdb.Connection,
}

type mediaService struct {
	repository scdb.RepositoryInterface
}

//保存对象,参数 : entity data.EntityInterface
func (this *mediaService) Save(entity data.EntityInterface) (sql.Result, error) {
	return this.repository.Save(entity)
}

//修改对象,参数 : entity data.EntityInterface
func (this *mediaService) Update(entity data.EntityInterface) (sql.Result, error) {
	return this.repository.Update(entity)
}

//保存或修改对象,参数 : entity data.EntityInterface
func (this *mediaService) SaveOrUpdate(entity data.EntityInterface) (sql.Result, error) {
	return this.repository.SaveOrUpdate(entity)
}

//查询多条,参数 : entity data.EntityBeanInterface
func (this *mediaService) Select(bean data.EntityBeanInterface) error {
	return this.repository.Select(bean)
}

//分页查询,参数 : entity data.EntityBeanInterface
func (this *mediaService) SelectPage(entityBean data.EntityBeanInterface, page *data.Page) error{
	return this.repository.SelectPage(entityBean, page)
}

//分页数量,参数 : entity data.EntityInterface
func (this *mediaService) SelectCount(entity data.EntityInterface) (int, error){
	return this.repository.SelectCount(entity)
}

//查询一条,参数 : entity data.EntityInterface
func (this *mediaService) SelectOne(entity data.EntityInterface) error {
	return this.repository.SelectOne(entity)
}

//删除,参数 : entity data.EntityInterface
func (this *mediaService) Delete(entity data.EntityInterface) (sql.Result, error) {
	return this.repository.Delete(entity)
}

//执行自定义DML语言. (DDL,DCL待添加)
func (this *mediaService) Execute(sql string, args ...interface{}) (sql.Result, error) {
	return this.repository.Execute(sql, args...)
}

//执行自定义DML语言. (DDL,DCL待添加)
func (this *mediaService) Query(sql string, args ...interface{}) (data.QueryResult, error) {
	return this.repository.Query(sql, args...)
}

//map保存
func (this *mediaService) SaveForMap(table string, data map[string][]string) (sql.Result, error) {
	return this.repository.SaveForMap(table, data)
}
