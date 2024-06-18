package influx

import (
	"LogCollector/logic/consts"
	"LogCollector/logic/model"
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

// WriteDiskInfo 磁盘信息
func WriteDiskInfo() {
	usageMap := make(map[string]*disk.UsageStat, 16)
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get disk info failed, err:%v", err)
	}
	for _, part := range partitions { // 存储所有分区
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			fmt.Printf("get disk info failed, err:%v", err)
			return
		}
		fmt.Printf("disk info:%v\n", usage)
		usageMap[usage.Path] = usage
	}

	writePoints(model.SystemInfo{
		InfoType: consts.DISK_INFO_TYPE,
		Data:     model.DiskInfo{PartitionUsageStat: usageMap},
	})
}
