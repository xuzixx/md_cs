package models

import (
	"github.com/astaxie/beego/orm"

	//
	_ "github.com/mattn/go-sqlite3"
)

// Chapter ...
type Chapter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Book *Book `orm:"rel(fk)" json:"-"`
}

// AddOne obj 是指针
func AddOne(obj interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(obj)
}

// GetOne obj 是指针
func GetOne(obj interface{}, cols ...string) error {
	o := orm.NewOrm()
	return o.Read(obj, cols...)
}

// GetAll obj 是指针
func GetAll(ptrStructOrTableName, objs interface{},
	limit, offset int, cols ...string) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("book")
	return qs.Limit(limit, offset).All(objs, cols...)
}

func init() {
	orm.RegisterDataBase("default", "sqlite3", "data/data.db")
	orm.RegisterModel(new(Book), new(Chapter))
	orm.RunSyncdb("default", false, true)
}
