package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	// 1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区
	config.Producer.Return.Successes = true                   // 成功发送的消息将被返回

	// 2.连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("produce closed, err:", err)
		return
	}
	defer client.Close()

	// 3.封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test_log"
	msg.Value = sarama.StringEncoder("This is a test log.")

	// 4.发送消息
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:", err)
		return
	}
	fmt.Printf("send message success, partition: %d, offset: %d\n", partition, offset)
}
