package models

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
	return
}
