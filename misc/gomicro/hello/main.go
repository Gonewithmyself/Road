package main

import (
	"hello/handler"
	"hello/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	hello "hello/proto/hello"
)

var etcdReg registry.Registry

func init() {
	etcdReg = etcd.NewRegistry(
		registry.Addrs("127.0.0.1:23791"),
	)
}

func main() {
	registry.NewRegistry()
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.hello"),
		micro.Version("latest"),
		micro.Registry(etcdReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.foo.service.hello", service.Server(), new(subscriber.Hello))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
