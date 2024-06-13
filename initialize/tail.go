package initialize

import (
	"LogCollector/logic/model"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func InitTail(collectConfigs []model.CollectConfig, MsgChan chan *sarama.ProducerMessage) {
	// 每一个配置，创建一个tail
	for _, cfg := range collectConfigs {

		// 创建日志收集任务
		task, err := model.NewTailTask(cfg.Path, cfg.Topic)
		if err != nil {
			logrus.Error("tail file: ===>", cfg.Path, "<=== failed, err:", err)
		}

		go task.Run(MsgChan)
	}

	return
}
