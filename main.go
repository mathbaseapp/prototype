package main

import (
	"flag"

	"prototype.mathbase.app/controller"
	"prototype.mathbase.app/service"
)

func main() {
	mode := flag.String("mode", "server", "represents mode in which program will be started. server or crawle is available. default is server.")
	flag.Parse()

	switch *mode {
	case "server":
		controller.RunServer()
	case "crawle":
		service.Crawle()
	default:
		panic("the argument of mode is invalid. only server or crawl are allowed.")
	}
}
