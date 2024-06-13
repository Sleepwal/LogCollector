package initialize

import (
	"LogCollector/config"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

// InitKafka
// @Description: 初始化kafka连接
// @param kafka_config
// @return *sarama.SyncProducer
func InitKafka(kafkaConfig config.KafKaConfig) sarama.SyncProducer {
	// 1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区
	config.Producer.Return.Successes = true                   // 成功发送的消息将被返回

	// 2.连接kafka
	client, err := sarama.NewSyncProducer(kafkaConfig.Address, config)
	if err != nil {
		logrus.Info("produce closed, err:", err)
		return nil
	}

	logrus.Info("kafka init success!")
	return client
}
