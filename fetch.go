package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/astaxie/beego/httplib"
	"github.com/tidwall/gjson"
	"github.com/xuzixx/md_cs/g"
)

//
var (
	MAXPARALLEL = 20
	sema        = make(chan struct{}, MAXPARALLEL)
)

func parseIndexHTML(content string) (map[string]string, error) {
	idTitles := make(map[string]string)

	for _, matchItem := range g.REChapter().FindAllStringSubmatch(content, -1) {
		id := matchItem[1]
		title := matchItem[2]

		if title != "开始阅读" && title != "最新章节" {
			idTitles[id] = title
		}
	}

	return idTitles, nil
}

// FetchOmnibus ...
func FetchOmnibus(b g.BookConfig) (map[string]string, error) {
	req := httplib.Get(b.BookURL)
	str, err := req.String()
	if err != nil {
		return nil, err
	}
	return parseIndexHTML(str)
}

func parseConfigJSON(content string) ([]string, error) {
	pagePostfix := []string{}
	jTitle := gjson.Get(content, "meta.title").String()
	jPages := gjson.Get(content, "pages.page").Map()

	for _, v := range jPages {
		pagePostfix = append(pagePostfix, v.String())
	}

	if len(pagePostfix) == 0 {
		err := fmt.Errorf("Title: %s has no pages in conf", jTitle)
		return pagePostfix, err
	}
	return pagePostfix, nil
}

// FetchChapter ...
func FetchChapter(id string) ([]string, error) {
	configURL := fmt.Sprintf(g.Config().URLConfTmp, id)

	req := httplib.Get(configURL)
	str, err := req.String()
	if err != nil {
		return nil, err
	}
	return parseConfigJSON(str)
}

// FetchImg ...
func FetchImg(id string, path string, postfix string, wg *sync.WaitGroup) {
	sema <- struct{}{}
	defer func() { <-sema }()
	defer wg.Done()

	imgURL := fmt.Sprintf(g.Config().URLImgTmp, id, postfix)
	referURL := fmt.Sprintf(g.Config().HeaderReferTmp, id)

	items := strings.Split(postfix, "/")
	filename := items[len(items)-1]
	filePth := fmt.Sprintf("%s/%s", path, filename)

	req := httplib.Get(imgURL)
	req.Header("Origin", g.Config().HeaderOrigin)
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/10.0.1 Safari/602.2.14")
	req.Header("Referer", referURL)
	err := req.ToFile(filePth)
	if err != nil {
		log.Printf("Error FetchImg %s to %s, %s\n", imgURL, path, err.Error())
	}
}

// FetchImgByURL ...
func FetchImgByURL(imgURL, path string) {
	items := strings.Split(imgURL, "/")
	id := items[4]
	log.Printf("id: %s\n", id)
	referURL := fmt.Sprintf(g.Config().HeaderReferTmp, id)

	filename := items[len(items)-1]
	filePth := fmt.Sprintf("%s/%s", path, filename)

	req := httplib.Get(imgURL)
	req.Header("Origin", g.Config().HeaderOrigin)
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/10.0.1 Safari/602.2.14")
	req.Header("Referer", referURL)
	err := req.ToFile(filePth)
	if err != nil {
		log.Printf("Error FetchImgByURL %s to %s, %s\n", imgURL, path, err.Error())
	}
}
