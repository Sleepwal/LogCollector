package main

import (
	"LogCollector/global"
	"LogCollector/logic"
	"LogCollector/logic/etcd"
	"LogCollector/logic/kafka"
	"LogCollector/logic/tail"
	"github.com/sirupsen/logrus"
)

func main() {
	// 1.初始化kafka连接
	err := kafka.InitKafka(global.CONFIG.KafKaConfig)
	if err != nil {
		global.LOG.Error(err)
		return
	}
	defer kafka.KAFKA_CONNECTOR.Close()
	global.LOG.Info("初始化kafka连接, ", kafka.KAFKA_CONNECTOR)

	// 2.初始化etcd连接
	err = etcd.InitEtcd(global.CONFIG.EtcdConfig.Endpoints)
	if err != nil {
		global.LOG.Error(err)
	}

	// 3.初始化tail
	collectConfigs, err := etcd.GetConfig(global.CONFIG.EtcdConfig.CollectKey)
	if err != nil {
		global.LOG.Fatalf("get config from etcd failed, err:%v", err)
		return
	}
	tail.InitTail(collectConfigs)

	err = logic.Run()
	if err != nil {
		logrus.Info("server run error: ", err)
		return
	}
}
