package main

import (
	"context"
	"fmt"
	"log"
	"misc/etcd/service"
	"os"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	endpoints = []string{
		// "localhost:23791",
		// "localhost:23792",
		"localhost:23793",
	}

	prefix = "/"
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
	// agent.WatchNodes()
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*40)
	rsp, er := agent.Client.Grant(ctx, 50)
	if er != nil {
		panic(er)
	}

	go func() {
		for i := 0; ; i++ {
			_, er = agent.Client.Put(ctx, "/go2/"+strconv.Itoa(i), "x", clientv3.WithLease(rsp.ID))
			if er != nil {
				log.Println(er)
				continue
			}
			if i+1%100 == 0 {
				time.Sleep(time.Millisecond * 20)
			}

			if i > 10000 {
				break
			}
		}
	}()

	for i := 0; ; i++ {
		_, er = agent.Client.Put(ctx, "/go/"+strconv.Itoa(i), "x", clientv3.WithLease(rsp.ID))
		if er != nil {
			log.Println(er)
			continue
		}
		if i+1%100 == 0 {
			time.Sleep(time.Millisecond * 20)
		}

		if i > 10000 {
			break
		}
	}

	_, er = agent.Client.Put(ctx, "/go2", "x", clientv3.WithLease(rsp.ID))
	if er != nil {
		panic(er)
	}

	_, er = agent.Client.Put(ctx, "/go22", "x", clientv3.WithLease(rsp.ID))
	if er != nil {
		panic(er)
	}

	lease2, er := agent.Client.Grant(ctx, 5)
	if er != nil {
		panic(er)
	}

	_, er = agent.Client.Put(ctx, "/go2", "x", clientv3.WithLease(lease2.ID))
	if er != nil {
		panic(er)
	}

	lkr, er := agent.Client.KeepAlive(context.Background(), lease2.ID)
	if er != nil {
		panic(er)
	}
	for k := range lkr {
		fmt.Println(k)
	}
	fmt.Println("done")
}
