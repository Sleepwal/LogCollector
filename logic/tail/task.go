package tail

import (
	"LogCollector/logic/kafka"
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type TailTask struct {
	Path   string `json:"path"`
	Topic  string `json:"topic"`
	TObj   *tail.Tail
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewTailTask(path, topic string) (task *TailTask, err error) {
	// 创建tail对象
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	tailFile, err := tail.TailFile(path, config)
	if err != nil {
		logrus.Error("create tail file: ", path, " failed, error:", err)
		return
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	task = &TailTask{
		Path:   path,
		Topic:  topic,
		TObj:   tailFile,
		Ctx:    ctx,
		Cancel: cancelFunc,
	}

	return
}

func (t *TailTask) Run() {
	// 循环读取日志
	logrus.Info("start tail file: ", t.Path)
	for {
		select {
		case <-t.Ctx.Done(): // 判断是否取消
			logrus.Info("tail file: ", t.Path, " context cancel, exit")
			return
		default:
			line, ok := <-t.TObj.Lines // 从tail的Lines通道中读取日志文件的数据
			if !ok {
				logrus.Warn("tail file closed, reopen ", t.Path)
				time.Sleep(time.Second) // 读取失败，等待1秒钟重新打开
				continue
			}
			if len(strings.Trim(line.Text, "\r")) == 0 { // 空行跳过
				continue
			}

			// 把一行日志封装成消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = t.Topic
			msg.Value = sarama.StringEncoder(line.Text)
			// 发送消息到通道
			kafka.MSG_CHAN <- msg
			fmt.Println("message: ", line.Text)
		}
	}
}
