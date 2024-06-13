package kafka

import (
	"LogCollector/config"
	"LogCollector/global"
	"errors"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

var (
	KAFKA_CONNECTOR sarama.SyncProducer          // kafka连接器
	MSG_CHAN        chan *sarama.ProducerMessage // kafka消息管道
)

// InitKafka
// @Description: 初始化kafka连接
// @param kafkaConfig
// @return err
func InitKafka(kafkaConfig config.KafKaConfig) (err error) {
	// 1.生产者配置
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	conf.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区
	conf.Producer.Return.Successes = true                   // 成功发送的消息将被返回

	// 2.连接kafka
	KAFKA_CONNECTOR, err = sarama.NewSyncProducer(kafkaConfig.Address, conf)
	if err != nil {
		return errors.New("produce closed, err: " + err.Error())
	}

	// 3.初始化消息管道
	MSG_CHAN = make(chan *sarama.ProducerMessage, global.CONFIG.KafKaConfig.ChanSize)

	logrus.Info("kafka init success!")
	return
}
