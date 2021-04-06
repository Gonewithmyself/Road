package main

import (
	"net/http"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/server"
	httpServer "github.com/micro/go-plugins/server/http/v2"
	"github.com/micro/go-plugins/transport/tcp/v2"
)

var etcdReg registry.Registry

func init() {
	etcdReg = etcd.NewRegistry(
		registry.Addrs("127.0.0.1:23791"),
	)

}

func main() {
	srv := httpServer.NewServer(
		server.Name("helloworld"),
		server.Registry(etcdReg),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})

	hd := srv.NewHandler(mux)

	srv.Handle(hd)

	registry.NewRegistry()
	// New Service
	service := micro.NewService(
		// micro.Name("com.foo.service.hello"),
		micro.Version("latest"),
		micro.Registry(etcdReg),
		micro.Transport(tcp.NewTransport()),
		micro.Server(srv),
	)

	// Initialise service
	service.Init()

	// // Register Handler
	// hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// // Register Struct as Subscriber
	// micro.RegisterSubscriber("com.foo.service.hello", service.Server(), new(subscriber.Hello))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
