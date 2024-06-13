package kafka

import "LogCollector/global"

// SendMsg
// @Description: 监听MSG_CHAN，发送消息到Kafka
// @param msg
// @return error
func SendMsg() {
	for {
		select {
		case msg := <-global.MSG_CHAN:
			partition, offset, err := global.KAFKA_CONNECTOR.SendMessage(msg)
			if err != nil {
				global.LOG.Warn("send message failed, err: ", err)
				return
			}
			global.LOG.Info("send message success, partition: ", partition, ", offset: ", offset)
		}
	}
}
