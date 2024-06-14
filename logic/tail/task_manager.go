package tail

import (
	"LogCollector/global"
	"LogCollector/logic/model"
)

type TailTaskMgr struct {
	TailTasks      map[string]*TailTask       // 存储所有的tail任务
	CollectConfigs []model.CollectConfig      // 存储所有的采集配置
	ConfigChan     chan []model.CollectConfig // 新配置通道
}

var (
	taskMgr *TailTaskMgr
)

// Watch
// @Description: 处理配置变化
// @receiver t
func (t *TailTaskMgr) Watch() {
	newConfig := <-t.ConfigChan
	//t.CollectConfigs = newConfig
	global.LOG.Info("tail任务配置有更新: ", newConfig)

	// 1.停止不需要的tail任务
	for key, task := range t.TailTasks {
		found := false
		for _, cfg := range newConfig { // 查找在新配置中不存在的任务
			if key == cfg.Path {
				found = true
				break
			}
		}
		if !found {
			task.Cancel()            // 停止tail任务
			delete(t.TailTasks, key) // 在map中删除
			global.LOG.Info("停止tail任务: ", key)
		}
	}

	for _, cfg := range newConfig {
		// 2.存在配置并运行中的tail任务不用管
		if t.isExist(cfg) {
			continue
		}
		// 3.启动新的tail任务
		task, err := NewTailTask(cfg.Path, cfg.Topic)
		if err != nil {
			global.LOG.Error("tail file: ===>", cfg.Path, "<=== failed, err:", err)
			continue
		}

		global.LOG.Info("启动新tail任务: ", task.Path)
		taskMgr.TailTasks[task.Path] = task // 存储tailTask任务
		go task.Run()
	}
}

func (t *TailTaskMgr) isExist(config model.CollectConfig) bool {
	_, ok := t.TailTasks[config.Path]
	return ok
}

func SendNewConfig(newConfigs []model.CollectConfig) {
	taskMgr.ConfigChan <- newConfigs
}
