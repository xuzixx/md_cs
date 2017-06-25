package models

import (
	"github.com/astaxie/beego/orm"

	//
	_ "github.com/mattn/go-sqlite3"
)

// Book ...
type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Chapters []*Chapter `orm:"reverse(many)" json:"-"`
}

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

func init() {
	orm.RegisterDataBase("default", "sqlite3", "data/data.db")
	orm.RegisterModel(new(Book), new(Chapter))
	orm.RunSyncdb("default", false, true)
}
