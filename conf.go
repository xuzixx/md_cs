package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type BookConfig struct {
	BookName string `json:"c_book"`
	BookUrl  string `json:"c_url"`
}

type Config struct {
	ResultBasePath string `json:"result_base_path"`

	URLConfTmp string `json:"url_conf_tmp"`
	URLImgTmp  string `json:"url_img_tmp"`

	HeaderReferTmp string `json:"header_refer_tmp"`
	HeaderOrigin   string `json:"header_origin"`

	REChapter string `json:"re_chapter"`

	Books []BookConfig `json:"books"`
}

func loadConfig(filePth string) (Config, error) {
	config := Config{}
	if _, err := os.Stat(filePth); os.IsNotExist(err) {
		return config, err
	}

	log.Printf("Loading config %s ...\n", filePth)
	text, _ := ioutil.ReadFile(filePth)
	err := json.Unmarshal(text, &config)

	RE_CHAPTER = regexp.MustCompile(config.REChapter)
	// TODO check necessary key

	return config, err
}
