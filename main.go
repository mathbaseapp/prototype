package main

import (
	"flag"

	"prototype.mathbase.app/controller"
	"prototype.mathbase.app/crawler"
)

func main() {
	// usage: go run main.go -mode=crawle
	mode := flag.String("mode", "server", "represents mode which program will be started in. server or crawl is available. default is server.")
	flag.Parse()

	switch *mode {
	case "server":
		controller.RunServer()
	case "crawle":
		token := "04955c64db710699566b3420e4a8ae01ec907dd6"
		c := crawler.NewCrawler(token, &crawler.QiitaArticleProcessor{})
		c.Crawle()
	default:
		panic("the argument of mode is invalid. only server or crawl are allowed.")
	}
}
