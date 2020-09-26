package main

import (
	"flag"
	"misc/tool/htmlviewer/router"
)

func main() {
	s := router.NewHttpServer(":" + port)
	s.Start()
}

var port string

func init() {
	flag.StringVar(&port, "p", "8080", "./xxx [-p port]")
	flag.Parse()
}
