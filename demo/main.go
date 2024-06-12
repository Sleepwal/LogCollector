package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		// 接收外部信号后，实现退出
		select {
		case <-ctx.Done():
			fmt.Println("Received exit signal")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("Done")
}
