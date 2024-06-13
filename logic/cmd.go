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

	// 循环读取日志
	//for {
	//	line, ok := <-global.TAIL.Lines // 从通道中读取日志（Lines chan *Line）
	//	if !ok {
	//		global.LOG.Warn("tail file closed, reopen ", global.CONFIG.FilePath)
	//		time.Sleep(time.Second) // 读取失败，等待1秒钟重新打开
	//		continue
	//	}
	//	if len(strings.Trim(line.Text, "\r")) == 0 { // 空行跳过
	//		continue
	//	}
	//
	//	// 把一行日志封装成消息
	//	msg := &sarama.ProducerMessage{}
	//	msg.Topic = global.CONFIG.Topic
	//	msg.Value = sarama.StringEncoder(line.Text)
	//	// 发送消息到通道
	//	global.MSG_CHAN <- msg
	//	fmt.Println("message: ", line.Text)
	//}
	select {}
	return nil
}
