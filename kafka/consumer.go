package kafka

import (
	"LogCollector/global"
	"github.com/IBM/sarama"
)

func ReceiveFromKafka(topic string) {
	kafkaCfg := global.CONFIG.KafKaConfig
	consumer, err := sarama.NewConsumer(kafkaCfg.Address, nil)
	if err != nil {
		global.LOG.Warn("fail to open kafka consumer, err:", err)
		return
	}

	partitions, err := consumer.Partitions(kafkaCfg.Topic)
	if err != nil {
		global.LOG.Warn("fail to get partitions, err:", err)
		return
	}

	// 遍历所有分区
	for partition := range partitions {
		// 每一个分区建立一个consumer
		pc, err := consumer.ConsumePartition(kafkaCfg.Topic, int32(partition), sarama.OffsetOldest)
		if err != nil {
			global.LOG.Warn("fail to start consume for partition, err:", err)
			return
		}
		defer pc.AsyncClose()

		// 获取分区中的消息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				global.LOG.Info("receive message from kafka, topic:", topic,
					", partition:", msg.Partition,
					", offset:", msg.Offset,
					", key:", string(msg.Key),
					", value:", string(msg.Value))
			}
		}(pc)

		select {}
	}
}
