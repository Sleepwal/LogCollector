package influx

import (
	"LogCollector/logic/consts"
	"LogCollector/logic/model"
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

// 内存信息
func WriteMemInfo() {
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v", err)
	}
	memInfo := model.MemInfo{
		Total:       info.Total,
		Available:   info.Available,
		Used:        info.Used,
		UsedPercent: info.UsedPercent,
		Buffers:     info.Buffers,
		Cached:      info.Cached,
	}
	fmt.Printf("memory into :%#v\n", memInfo)

	writePoints(model.SystemInfo{
		InfoType: consts.MEM_INFO_TYPE,
		Data:     memInfo,
	})
}
