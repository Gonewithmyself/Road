package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type Agent struct {
	Path   string
	Nodes  map[string]*Node
	Client *clientv3.Client
}

//node is a client
type Node struct {
	State bool
	Key   string
	Info  ServiceInfo
}

func NewAgent(endpoints []string, watchPath string) (*Agent, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second,
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	master := &Agent{
		Path:   watchPath,
		Nodes:  make(map[string]*Node),
		Client: cli,
	}

	return master, err
}

func (m *Agent) AddNode(key string, info *ServiceInfo) {
	node := &Node{
		State: true,
		Key:   key,
		Info:  *info,
	}

	m.Nodes[node.Key] = node
}

func GetServiceInfo(ev *clientv3.Event) *ServiceInfo {
	info := &ServiceInfo{}
	err := json.Unmarshal([]byte(ev.Kv.Value), info)
	if err != nil {
		log.Println(err)
	}
	return info
}

func (m *Agent) WatchNodes() {
	rch := m.Client.Watch(context.Background(), m.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				info := GetServiceInfo(ev)
				m.AddNode(string(ev.Kv.Key), info)
			case clientv3.EventTypeDelete:
				fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				delete(m.Nodes, string(ev.Kv.Key))
			}
		}
	}
}
