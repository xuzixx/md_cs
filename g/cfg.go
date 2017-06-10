package g

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
	"sync"

	"gosdk/utils/bdfile"
)

// BookConfig ...
type BookConfig struct {
	BookName string `json:"c_book"`
	BookURL  string `json:"c_url"`
}

// GlobalConfig ...
type GlobalConfig struct {
	ResultBasePath string `json:"result_base_path"`

	URLConfTmp string `json:"url_conf_tmp"`
	URLImgTmp  string `json:"url_img_tmp"`

	HeaderReferTmp string `json:"header_refer_tmp"`
	HeaderOrigin   string `json:"header_origin"`

	REChapter string `json:"re_chapter"`

	Books []BookConfig `json:"books"`
}

var (
	config     *GlobalConfig
	reChapter  *regexp.Regexp
	configLock = new(sync.RWMutex)
)

// Config ...
func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// REChapter ...
func REChapter() *regexp.Regexp {
	configLock.RLock()
	defer configLock.RUnlock()

	return reChapter
}

// ParseConfig ...
func ParseConfig(cfg string) {
	var c GlobalConfig

	if !bdfile.FileExists(cfg) {
		log.Fatalf("Config file %s not exists", cfg)
	}

	buf, err := ioutil.ReadFile(cfg)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err = json.Unmarshal(buf, &c); err != nil {
		log.Fatalf("Parse config file %s error: %s", cfg, err.Error())
	}

	configLock.Lock()
	defer configLock.Unlock()

	config = &c
	reChapter = regexp.MustCompile(c.REChapter)

	log.Printf("Parse config file: %s successfully", cfg)
}
