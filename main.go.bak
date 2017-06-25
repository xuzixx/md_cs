package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/urfave/cli"
	"github.com/xuzixx/md_cs/g"
)

//
var (
	wg sync.WaitGroup
)

func start() {
	for _, book := range g.Config().Books {
		log.Printf("Start Fetch book %s\n", book.BookName)
		idTitles, err := FetchOmnibus(book)
		if err != nil {
			log.Printf("Error FetchOmnibus %s: %s\n", book.BookName, err.Error())
			continue
		}

		for id, title := range idTitles {
			msg := fmt.Sprintf("%s-%s-%s", book.BookName, id, title)
			path := fmt.Sprintf("%s/%s/%s", g.Config().ResultBasePath, book.BookName, title)

			if _, err := os.Stat(path); os.IsNotExist(err) {
				err = os.MkdirAll(path, 0777)
				if err != nil {
					log.Printf("Error mkdir %s", path)
					continue
				}

				postfix, err := FetchChapter(id)
				if err != nil {
					log.Printf("Error %s, %s", msg, err.Error())
					continue
				}

				for _, p := range postfix {
					wg.Add(1)
					go FetchImg(id, path, p, &wg)
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
	app.Version = g.VERSION
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

		g.ParseConfig(confFile)

		start()
		return nil
	}

	app.Run(os.Args)
}
