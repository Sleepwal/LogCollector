package initialize

import (
	"LogCollector/logic/model"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var client *clientv3.Client

func InitEtcd(endpoints []string) *clientv3.Client {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logrus.Error("etcd connect failed, err:", err)
		return nil
	}

	client = etcdClient
	return etcdClient
}

func GetConfig(key string) (collects []model.CollectConfig, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.Get(ctx, key)
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
