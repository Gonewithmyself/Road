package main

import (
	"log"
	"strings"
	"time"

	nats "github.com/nats-io/nats.go"
)

var nc *nats.Conn

func main() {
	servers := []string{"nats://127.0.0.1:5220", "nats://127.0.0.1:5221", "nats://127.0.0.1:5222"}

	var err error
	nc, err = nats.Connect(strings.Join(servers, ", "))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	log.Println("connect sussess")
	// pubsub()

	repResp()
	time.Sleep(time.Second)
}

func repResp() {
	log.Println("------------------------------------")

	subapply("go")

	request("go")

}

func pubsub() {
	log.Println("------------------------------------")
	subAsync("test")
	subAsync("test.*")
	subAsync("test.*.*")
	subAsync("test.*.go")
	subAsync("test.>")

	publish()
}

func subAsync(key string) {
	nc.Subscribe(key, func(m *nats.Msg) {
		log.Printf("recv key(%v)\t \t%v\n", key, m)
	})
}

func subapply(key string) {
	nc.Subscribe(key, func(m *nats.Msg) {
		log.Printf("request key(%v)\t \t%v\n", key, m)
		err := m.Respond([]byte("hello"))
		if err != nil {
			log.Println("rsp ", err)
		}

		log.Println("rsp ok")
	})
}

func publish() {
	nc.Publish("test", []byte("hello"))
	nc.Publish("test.123.321", []byte("hello"))
	nc.Publish("test.123.go", []byte("hello"))
	nc.Publish("test.python", []byte("hello"))
}

func request(key string) {
	m, er := nc.Request(key, []byte("999"), time.Second)
	if er != nil {
		log.Println("request", er)
		return
	}

	log.Println("reply", string(m.Data))
}
