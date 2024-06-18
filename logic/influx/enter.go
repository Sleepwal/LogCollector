package influx

import "time"

func WriteSystemInfo() {
	InfluxDbConn("Organization", "LogCollector")
	defer CLIENT.Close()

	go func() {
		for { // 定时写入
			WriteCpuInfo()
			WriteMemInfo()
			time.Sleep(time.Second * 10)
		}
	}()
}
