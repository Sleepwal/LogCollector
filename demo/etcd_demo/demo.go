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
		panic(err)
	}
	defer client.Close()

	// put a key-value pair
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//str := `[{"path":"./demo/web.log","topic":"web_log"}]`
	str := `[{"path":"./demo/web.log","topic":"web_log"},{"path":"./demo/redis.log","topic":"redis_log"}]`
	_, err = client.Put(ctx, "log_collector", str)
	if err != nil {
		fmt.Println("put failed:", err)
		return
	}
	cancel()

	// get the value of the key
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	get, err := client.Get(ctx, "log_collector")
	if err != nil {
		fmt.Println("get failed:", err)
		return
	}
	for _, v := range get.Kvs {
		fmt.Printf("key: %s, value: %s\n", string(v.Key), string(v.Value))
	}
	cancel()
}
