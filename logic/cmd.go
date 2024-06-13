package logic

import (
	"LogCollector/global"
	"LogCollector/logic/kafka"
)

// Run
// @Description: 业务逻辑
// @return error
func Run() error {
	go kafka.SendMsg()                             // 启动发送消息的协程
	go kafka.ReceiveFromKafka(global.CONFIG.Topic) // 启动从Kafka接收消息的协程

	select {}
}
