package tail

import (
	"LogCollector/logic/model"
	"github.com/sirupsen/logrus"
)

func InitTail(collectConfigs []model.CollectConfig) {
	taskMgr = &TailTaskMgr{
		TailTasks:      make(map[string]*TailTask, 20),
		CollectConfigs: collectConfigs,
		ConfigChan:     make(chan []model.CollectConfig),
	}

	// 每一个配置，创建一个tail
	for _, cfg := range collectConfigs {
		// 创建日志收集任务
		task, err := NewTailTask(cfg.Path, cfg.Topic)
		if err != nil {
			logrus.Error("tail file: ===>", cfg.Path, "<=== failed, err:", err)
			continue
		}

		// 存储tailTask任务
		taskMgr.TailTasks[task.Path] = task
		go task.Run()
	}

	// 监听，获取新配置
	go taskMgr.Watch()
	return
}
