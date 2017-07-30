package tools

import (
	"regexp"

	"github.com/astaxie/beego/httplib"
)

func parseIndexHTML(content, re string) map[string]string {
	idTitles := make(map[string]string)

	//beego.AppConfig.String("smp:re-chapter")
	for _, matchItem := range regexp.MustCompile(re).FindAllStringSubmatch(content, -1) {
		id := matchItem[1]
		title := matchItem[2]

		if title != "开始阅读" && title != "最新章节" {
			idTitles[id] = title
		}
	}

	return idTitles
}

// FetchOmnibus 第一步, 根据book url 抓取id + title
// Return: map: key: id, value: title
func FetchOmnibus(url, re string) (map[string]string, error) {
	req := httplib.Get(url)
	str, err := req.String()
	if err != nil {
		return nil, err
	}

	return parseIndexHTML(str, re), nil
}
