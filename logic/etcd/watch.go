package etcd

import (
	"LogCollector/global"
	"LogCollector/logic/model"
	"LogCollector/logic/tail"
	"context"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// WatchConfig
// @Description: 监控etcd配置变化
// @param key
func WatchConfig(key string) {
	for {
		watchChan := ETCD_CLIENT.Watch(context.Background(), key)
		// 获取最新的配置
		for watchResp := range watchChan {
			for _, evt := range watchResp.Events {
				global.LOG.Infof("Type: %s, Key: %s, Value: %s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				var newConfig []model.CollectConfig
				if evt.Type == clientv3.EventTypeDelete { // 删除操作
					global.LOG.Warnf("etcd delete the key\n")
					tail.SendNewConfig(newConfig)
					continue
				}

				err := json.Unmarshal(evt.Kv.Value, &newConfig)
				if err != nil {
					global.LOG.Errorf("unmarshal config error: %v\n", err)
				}

				// tail中使用新的配置
				tail.SendNewConfig(newConfig)
			}
		}
	}
}
