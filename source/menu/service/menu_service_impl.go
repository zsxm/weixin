//scgen
package service

import (
	"database/sql"
	"github.com/zsxm/scgo/data"
	"github.com/zsxm/scgo/data/scdb"
)

var MenuService = menuService{
	repository: scdb.Connection,
}

type menuService struct {
	repository scdb.RepositoryInterface
}

//保存对象,参数 : entity data.EntityInterface
func (this *menuService) Save(entity data.EntityInterface) (sql.Result, error) {
	return this.repository.Save(entity)
}

//修改对象,参数 : entity data.EntityInterface
func (this *menuService) Update(entity data.EntityInterface) (sql.Result, error) {
	return this.repository.Update(entity)
}

//保存或修改对象,参数 : entity data.EntityInterface
func (this *menuService) SaveOrUpdate(entity data.EntityInterface) (sql.Result, error) {
	return this.repository.SaveOrUpdate(entity)
}

//查询多条,参数 : entity data.EntityBeanInterface
func (this *menuService) Select(bean data.EntityBeanInterface) error {
	return this.repository.Select(bean)
}

//分页查询,参数 : entity data.EntityBeanInterface
func (this *menuService) SelectPage(entityBean data.EntityBeanInterface, page *data.Page) error{
	return this.repository.SelectPage(entityBean, page)
}

//分页数量,参数 : entity data.EntityInterface
func (this *menuService) SelectCount(entity data.EntityInterface) (int, error){
	return this.repository.SelectCount(entity)
}

//查询一条,参数 : entity data.EntityInterface
func (this *menuService) SelectOne(entity data.EntityInterface) error {
	return this.repository.SelectOne(entity)
}

//删除,参数 : entity data.EntityInterface
func (this *menuService) Delete(entity data.EntityInterface) (sql.Result, error) {
	return this.repository.Delete(entity)
}

//执行自定义DML语言. (DDL,DCL待添加)
func (this *menuService) Execute(sql string, args ...interface{}) (sql.Result, error) {
	return this.repository.Execute(sql, args...)
}

//执行自定义DML语言. (DDL,DCL待添加)
func (this *menuService) Query(sql string, args ...interface{}) (data.QueryResult, error) {
	return this.repository.Query(sql, args...)
}

//map保存
func (this *menuService) SaveForMap(table string, data map[string][]string) (sql.Result, error) {
	return this.repository.SaveForMap(table, data)
}
