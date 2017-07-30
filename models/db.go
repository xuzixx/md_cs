package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/xuzixx/md_cs/tools"
)

// Book ...
type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	URL       string `json:"url"`

	Chapters []*Chapter `orm:"reverse(many)" json:"-"`
}

// Chapter ...
type Chapter struct {
	ID    int    `json:"id"`
	URLID string `json:"url_id"`
	Name  string `json:"name"`

	Book *Book `orm:"rel(fk)" json:"-"`
}

// Message ...
type Message struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Start TODO
func (b *Book) Start() (err error) {
	// 第一步, 抓出来所有章节
	// 第二步, 校验所有章节是否存在于数据库
	//     全量对比每个章节的id, title, url
	//     得到所有没抓过的
	// 第三步, 入一个消息库
	// 第四步, 没有的重新抓
	if b.Completed {
		return
	}

	re := beego.AppConfig.String("smp:ReChapter")
	basePath := beego.AppConfig.String("ResultBasePath")
	idTitles, err := tools.FetchOmnibus(b.URL, re)
	if err != nil {
		return
	}

	for id, title := range idTitles {
		c := &Chapter{URLID: id}
		err = GetOne(c, "URLID")
		if err != nil {
			logs.Critical("")
		}
		//NOW TODO

		path := fmt.Sprintf("%s/%s/%s", basePath, b.Name, title)
		if _, err = os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, 0777)
			if err != nil {
				return
			}

		}

	}

	return
}
