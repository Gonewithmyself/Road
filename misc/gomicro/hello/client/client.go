package main

import (
	"context"
	"fmt"
	hello "hello/proto/hello"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/transport/tcp/v2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.hello.client"), //name the client service
		micro.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:23791"))),
		micro.Transport(tcp.NewTransport()),
	)
	// Initialise service
	service.Init()

	//create hello service client
	helloClient := hello.NewHelloService("com.foo.service.hello", service.Client())

	//invoke hello service method
	resp, err := helloClient.Call(context.TODO(), &hello.Request{Name: "Bill 4"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Msg)
}
