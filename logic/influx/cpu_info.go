package influx

import (
	"LogCollector/logic/consts"
	"LogCollector/logic/model"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"time"
)

func WriteCpuInfo() {
	// cpu使用率
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)

	// 写入influxdb
	writePoints(model.SystemInfo{
		InfoType: consts.CPU_INFO_TYPE,
		Data:     model.CpuInfo{CpuPercent: percent[0]},
	})
}

// WriteCpuLoad
// @Description: 向influxdb写入cpu负载
func WriteCpuLoad() {
	avg, err := load.Avg()
	if err != nil {
		fmt.Printf("get cpu load failed, err:%v", err)
	}
	fmt.Printf("cpu load:%v\n", avg)
}
