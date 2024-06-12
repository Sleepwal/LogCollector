package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	fileName := "../test.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}

	// 打开文件
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file error:", err)
		return
	}

	// 循环读取日志
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines // 从通道中读取日志（Lines chan *Line）
		if !ok {
			fmt.Println("tail file closed, reopen ", fileName)
			time.Sleep(time.Second) // 读取失败，等待1秒钟重新打开
			continue
		}
		fmt.Println("message: ", msg.Text)
	}
}
