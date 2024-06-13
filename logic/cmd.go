package logic

import (
	"LogCollector/global"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

// Run
// @Description: 业务逻辑
// @return error
func Run() error {
	go sendMsg() // 启动发送消息的协程

	// 循环读取日志
	for {
		line, ok := <-global.TAIL.Lines // 从通道中读取日志（Lines chan *Line）
		if !ok {
			global.LOG.Warn("tail file closed, reopen ", global.CONFIG.FilePath)
			time.Sleep(time.Second) // 读取失败，等待1秒钟重新打开
			continue
		}

		// 把一行日志封装成消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = global.CONFIG.Topic
		msg.Value = sarama.StringEncoder(line.Text)
		// 发送消息到通道
		global.MsgChan <- msg
		fmt.Println("message: ", line.Text)
	}
}

// sendMsg
// @Description: 监听msgChan，发送消息到Kafka
// @param msg
// @return error
func sendMsg() {
	for {
		select {
		case msg := <-global.MsgChan:
			partition, offset, err := global.KAFKA_CONNECTOR.SendMessage(msg)
			if err != nil {
				global.LOG.Warn("send message failed, err: ", err)
				return
			}
			global.LOG.Info("send message success, partition: ", partition, ", offset: ", offset)
		}
	}
}
