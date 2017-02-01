package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/urfave/cli"
)

var (
	RE_CHAPTER *regexp.Regexp
	wg         sync.WaitGroup
)

func start(c Config) {
	for _, book := range c.Books {
		log.Printf("Start Fetch book %s\n", book.BookName)
		idTitles, err := FetchOmnibus(book)
		if err != nil {
			log.Printf("Error FetchOmnibus %s: %s\n", book.BookName, err.Error())
			continue
		}

		for id, title := range idTitles {
			msg := fmt.Sprintf("%s-%s-%s", book.BookName, id, title)
			path := fmt.Sprintf("%s/%s/%s", c.ResultBasePath, book.BookName, title)

			if _, err := os.Stat(path); os.IsNotExist(err) {
				err = os.MkdirAll(path, 0777)
				if err != nil {
					log.Printf("Error mkdir %s", path)
					continue
				}

				postfix, err := FetchChapter(id, c)
				if err != nil {
					log.Printf("Error %s, %s", msg, err.Error())
					continue
				}

				for _, p := range postfix {
					wg.Add(1)
					go FetchImg(id, path, p, c, &wg)
				}
			} else {
				log.Printf("%s already exist", msg)
			}
		}
	}
	wg.Wait()
}

func main() {

	app := cli.NewApp()
	app.Name = "MD_CS"
	app.Version = "0.1"
	app.Usage = "Comic Spider"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf,c",
			Value: "config.json",
			Usage: "Specify Config file",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		confFile := ctx.String("conf")
		if _, err := os.Stat(confFile); os.IsNotExist(err) {
			log.Fatalf("config file does not exist: %s", confFile)
		}
		cnf, err := loadConfig(confFile)
		if err != nil {
			log.Fatalf("Failed to parse json config file: %s", confFile)
		}
		start(cnf)
		return nil
	}

	app.Run(os.Args)
}
