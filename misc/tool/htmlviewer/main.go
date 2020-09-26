package main

import "misc/tool/htmlviewer/router"

func main() {
	s := router.NewHttpServer(":8080")
	s.Start()
}
