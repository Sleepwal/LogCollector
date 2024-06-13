package global

import (
	"LogCollector/initial"
	"github.com/IBM/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

var (
	LOG             *logrus.Logger               // 日志
	CONFIG          *initial.SystemConfig        // 系统配置
	TAIL            *tail.Tail                   // 读取日志工具
	KAFKA_CONNECTOR sarama.SyncProducer          // kafka连接器
	MsgChan         chan *sarama.ProducerMessage // kafka消息管道
)

func init() {
	LOG = initial.InitLogger()
	CONFIG = initial.InitSystemConfig()
	TAIL = initial.InitTail(CONFIG.FilePath)
	KAFKA_CONNECTOR = initial.InitKafka(CONFIG.KafKaConfig)
	MsgChan = make(chan *sarama.ProducerMessage, CONFIG.KafKaConfig.ChanSize)
}
