package main

// import "prototype.mathbase.app/crawler"
import (
	"flag"

	"prototype.mathbase.app/crawler"

	"prototype.mathbase.app/controller"
)

func main() {
	// usage: go run main.go -mode=crawl
	mode := flag.String("mode", "server", "represents mode which program will be started in. server or crawl is available. default is server.")
	flag.Parse()

	switch *mode {
	case "server":
		controller.RunServer()
	case "crawl":
		crawler.Crawle()
	default:
		panic("the argument of mode is invalid. only server or crawl are allowed.")
	}
}
