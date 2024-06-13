package global

import (
	"LogCollector/config"
	"LogCollector/initialize"
	"github.com/IBM/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	LOG             *logrus.Logger               // 日志
	CONFIG          *config.SystemConfig         // 系统配置
	TAIL            *tail.Tail                   // 读取日志工具
	KAFKA_CONNECTOR sarama.SyncProducer          // kafka连接器
	MSG_CHAN        chan *sarama.ProducerMessage // kafka消息管道
	ETCD_CLIENT     *clientv3.Client             // etcd连接器
)

func init() {
	LOG = initialize.Logger()
	CONFIG = initialize.InitSystemConfig()
	//TAIL = initialize.InitTail(CONFIG.FilePath)
	KAFKA_CONNECTOR = initialize.InitKafka(CONFIG.KafKaConfig)
	MSG_CHAN = make(chan *sarama.ProducerMessage, CONFIG.KafKaConfig.ChanSize)
	ETCD_CLIENT = initialize.InitEtcd(CONFIG.EtcdConfig.Endpoints)
	// 从etcd中获取动态配置
	collectConfigs, err := initialize.GetConfig(CONFIG.EtcdConfig.CollectKey)
	if err != nil {
		LOG.Fatalf("get config from etcd failed, err:%v", err)
		return
	}
	initialize.InitTail(collectConfigs, MSG_CHAN)
	LOG.Infof("etcd动态配置: %v", collectConfigs)
}
