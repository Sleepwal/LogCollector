package initial

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

func InitTail(fileName string) *tail.Tail {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}

	// 打开文件
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		logrus.Error("create tail file: ", fileName, " failed, error:", err)
		return nil
	}

	logrus.Info("tail file: ===>", fileName, "<=== success")
	return tails
}
