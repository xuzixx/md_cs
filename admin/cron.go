package admin

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"github.com/xuzixx/md_cs/models"
)

// StartCS 定时任务抓去
func StartCS() (err error) {
	logs.Info("Start CS")
	var (
		count  int64
		limit  = 10
		offset = 0
	)

	count, err = fetchBooks(limit, offset)
	if err != nil {
		return
	}

	page := int(count) / limit
	for {
		if offset == page {
			break
		}
		offset++
		_, err = fetchBooks(limit, offset)
		if err != nil {
			return
		}
	}

	logs.Info("Start CS Done")
	return
}

func fetchBooks(limit, offset int) (count int64, err error) {
	var books []*models.Book
	count, err = models.GetAll(&models.Book{}, &books, limit, offset)
	if err != nil {
		return
	}

	for _, b := range books {
		err = b.Start()
		if err != nil {
			return
		}
	}

	return
}

// StartMSG 定时发送消息, TODO
func StartMSG() (err error) {
	return
}

func init() {
	tk := toolbox.NewTask("Daily check", "0 12 * * * *", StartCS)

	toolbox.AddTask(tk.Taskname, tk)
}
