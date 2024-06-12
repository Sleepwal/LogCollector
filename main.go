package main

import (
	"LogCollector/global"
	"LogCollector/logic"
	"github.com/sirupsen/logrus"
)

func main() {
	// 1.读取配置文件
	global.LOG.Info(global.CONFIG)

	// 2.初始化kafka连接
	global.LOG.Info("初始化kafka连接, ", global.KAFKA_CONNECTOR)
	defer global.KAFKA_CONNECTOR.Close()

	err := logic.Run()
	if err != nil {
		logrus.Info("server run error: ", err)
		return
	}
}
