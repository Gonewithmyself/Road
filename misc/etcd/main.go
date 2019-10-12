package main

import (
	"log"
	"misc/etcd/service"
	"os"
)

var (
	endpoints = []string{
		"localhost:32769",
		"localhost:32771",
		"localhost:32773",
	}

	prefix = "services/"
)

func main() {
	if len(os.Args) > 1 {
		runSvc(os.Args[1])
	} else {
		runAgent()
	}
}

func runSvc(name string) {
	svc, err := service.NewService(name, service.ServiceInfo{IP: "0.0.0"}, endpoints)
	if nil != err {
		log.Fatal("new service", err)
	}

	svc.Start()
}

func runAgent() {
	agent, err := service.NewAgent(endpoints, prefix)
	if nil != err {
		log.Fatal("new agent", err)
	}
	agent.WatchNodes()
}
