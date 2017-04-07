package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/astaxie/beego/httplib"
	"github.com/tidwall/gjson"
)

var (
	MAXPARALLEL = 20
	sema        = make(chan struct{}, MAXPARALLEL)
)

func parseIndexHTML(content string) (map[string]string, error) {
	idTitles := make(map[string]string)

	for _, matchItem := range RE_CHAPTER.FindAllStringSubmatch(content, -1) {
		id := matchItem[1]
		title := matchItem[2]

		if title != "开始阅读" && title != "最新章节" {
			idTitles[id] = title
		}
	}

	return idTitles, nil
}

func FetchOmnibus(b BookConfig) (map[string]string, error) {
	req := httplib.Get(b.BookUrl)
	str, err := req.String()
	if err != nil {
		return nil, err
	}
	return parseIndexHTML(str)
}

func parseConfigJson(content string) ([]string, error) {
	pagePostfix := []string{}
	jTitle := gjson.Get(content, "meta.title").String()
	jPages := gjson.Get(content, "pages.page").Map()

	for _, v := range jPages {
		pagePostfix = append(pagePostfix, v.String())
	}

	if len(pagePostfix) == 0 {
		err := errors.New(fmt.Sprintf("Title: %s has no pages in conf", jTitle))
		return pagePostfix, err
	}
	return pagePostfix, nil
}

func FetchChapter(id string, c Config) ([]string, error) {
	configURL := fmt.Sprintf(c.URLConfTmp, id)

	req := httplib.Get(configURL)
	str, err := req.String()
	if err != nil {
		return nil, err
	}
	return parseConfigJson(str)
}

func FetchImg(id string, path string, postfix string, c Config, wg *sync.WaitGroup) {
	sema <- struct{}{}
	defer func() { <-sema }()
	defer wg.Done()

	imgURL := fmt.Sprintf(c.URLImgTmp, id, postfix)
	referURL := fmt.Sprintf(c.HeaderReferTmp, id)

	items := strings.Split(postfix, "/")
	filename := items[len(items)-1]
	filePth := fmt.Sprintf("%s/%s", path, filename)

	req := httplib.Get(imgURL)
	req.Header("Origin", c.HeaderOrigin)
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/10.0.1 Safari/602.2.14")
	req.Header("Referer", referURL)
	err := req.ToFile(filePth)
	if err != nil {
		log.Printf("Error FetchImg %s to %s, %s\n", imgURL, path, err.Error())
	}
}

func FetchImgByURL(imgURL, path string, c Config) {
	items := strings.Split(imgURL, "/")
	id := items[4]
	log.Printf("id: %s\n", id)
	referURL := fmt.Sprintf(c.HeaderReferTmp, id)

	filename := items[len(items)-1]
	filePth := fmt.Sprintf("%s/%s", path, filename)

	req := httplib.Get(imgURL)
	req.Header("Origin", c.HeaderOrigin)
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/10.0.1 Safari/602.2.14")
	req.Header("Referer", referURL)
	err := req.ToFile(filePth)
	if err != nil {
		log.Printf("Error FetchImgByURL %s to %s, %s\n", imgURL, path, err.Error())
	}
}
