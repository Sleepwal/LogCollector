package main

import (
	"LogCollector/logic/influx"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"time"
)

var (
	write api.WriteAPIBlocking
)

func main() {
	influx.InfluxDbConn("Organization", "LogCollector")

	for {
		influx.WriteCpuInfo()
		influx.WriteMemInfo()
		influx.WriteDiskInfo()
		influx.WriteNetInfo()
		time.Sleep(time.Second)
	}
}
