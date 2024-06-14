package etcd

import (
	"LogCollector/global"
	"LogCollector/logic/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	ETCD_CLIENT *clientv3.Client // etcd连接器
)

func InitEtcd(endpoints []string) (err error) {
	ETCD_CLIENT, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return errors.New("etcd connect failed, err: " + err.Error())
	}

	go WatchConfig(global.CONFIG.EtcdConfig.CollectKey)

	return nil
}

func GetConfig(key string) (collects []model.CollectConfig, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := ETCD_CLIENT.Get(ctx, key)
	if err != nil {
		logrus.Errorf("etcd get key:%s failed, err:%v", key, err)
		return
	}
	if len(resp.Kvs) == 0 {
		logrus.Warnf("etcd get len=0, key:%s , err:%v", key, err)
		return
	}

	kv := resp.Kvs[0]

	// 反序列化
	err = json.Unmarshal(kv.Value, &collects)
	if err != nil {
		logrus.Errorf("json unmarshal failed, err:%v", err)
		return
	}

	return
}
