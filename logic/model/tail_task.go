package model

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type TailTask struct {
	path  string `json:"path"`
	topic string `json:"topic"`
	tObj  *tail.Tail
}

func NewTailTask(path, topic string) (task *TailTask, err error) {
	// 创建tail对象
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	tailFile, err := tail.TailFile(path, config)
	if err != nil {
		logrus.Error("create tail file: ", path, " failed, error:", err)
		return
	}

	task = &TailTask{
		path:  path,
		topic: topic,
		tObj:  tailFile,
	}

	return
}

func (t *TailTask) Run(MsgChan chan *sarama.ProducerMessage) {
	// 循环读取日志
	logrus.Info("start tail file: ", t.path)
	for {
		line, ok := <-t.tObj.Lines // 从通道中读取日志（Lines chan *Line）
		if !ok {
			logrus.Warn("tail file closed, reopen ", t.path)
			time.Sleep(time.Second) // 读取失败，等待1秒钟重新打开
			continue
		}
		if len(strings.Trim(line.Text, "\r")) == 0 { // 空行跳过
			continue
		}

		// 把一行日志封装成消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = t.topic
		msg.Value = sarama.StringEncoder(line.Text)
		// 发送消息到通道
		MsgChan <- msg
		fmt.Println("message: ", line.Text)
	}
}
