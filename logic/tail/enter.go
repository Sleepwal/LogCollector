package tail

import (
	"LogCollector/logic/model"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

var TAIL *tail.Tail // 读取日志工具

func InitTail(collectConfigs []model.CollectConfig) {
	// 每一个配置，创建一个tail
	for _, cfg := range collectConfigs {

		// 创建日志收集任务
		task, err := NewTailTask(cfg.Path, cfg.Topic)
		if err != nil {
			logrus.Error("tail file: ===>", cfg.Path, "<=== failed, err:", err)
			continue
		}

		go task.Run()
	}

	return
}
