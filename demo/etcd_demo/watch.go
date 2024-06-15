package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd failed:", err)
		return
	}
	defer client.Close()

	// watch the key
	watchResp := client.Watch(context.Background(), "name")
	for w := range watchResp {
		for _, event := range w.Events {
			fmt.Println("event:", event.Type, string(event.Kv.Key), string(event.Kv.Value))
		}
	}
}