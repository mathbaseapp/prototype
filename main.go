package main

import (
	"flag"

	"prototype.mathbase.app/controller"
	"prototype.mathbase.app/converter"
	"prototype.mathbase.app/crawler"
	"prototype.mathbase.app/tokenizer"
)

func main() {
	mode := flag.String("mode", "server", "represents mode in which program will be started. server or crawl is available. default is server.")
	flag.Parse()

	switch *mode {
	case "server":
		controller.RunServer()
	case "crawle":
		token := "04955c64db710699566b3420e4a8ae01ec907dd6"
		parser := converter.GetParser(converter.Latex)
		tokenizer := &tokenizer.MathmlTokenizer{}
		c := crawler.NewCrawler(token, &crawler.QiitaArticleProcessor{Parser: parser, Tokenizer: tokenizer})
		c.Crawle()
	default:
		panic("the argument of mode is invalid. only server or crawl are allowed.")
	}
}
