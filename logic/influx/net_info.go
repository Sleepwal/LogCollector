package influx

import (
	"LogCollector/logic/consts"
	"LogCollector/logic/model"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"time"
)

var (
	lastNetIOStatTimeStamp int64
	lastNetInfo            *model.NetInfo
)

// WriteNetInfo 网络信息
func WriteNetInfo() {
	netMap := make(map[string]*model.IOStat, 8)
	curTimeStamp := time.Now().Unix()
	counters, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net info failed, err:%v", err)
		return
	}
	for _, counter := range counters {
		stat := &model.IOStat{
			BytesSent:   counter.BytesSent,
			BytesRecv:   counter.BytesRecv,
			PacketsSent: counter.PacketsSent,
			PacketsRecv: counter.PacketsRecv,
		}
		// 添加到map中
		netMap[counter.Name] = stat

		if lastNetInfo == nil || lastNetIOStatTimeStamp == 0 {
			continue
		}
		// 计算网卡相关速率
		diff := float64(curTimeStamp - lastNetIOStatTimeStamp)
		stat.BytesSentRate = float64(stat.BytesSent-lastNetInfo.NetIOCountersStat[counter.Name].BytesSent) / diff
		stat.BytesRecveRate = float64(stat.BytesRecv-lastNetInfo.NetIOCountersStat[counter.Name].BytesRecv) / diff
		stat.PacketsSentRate = float64(stat.PacketsSent-lastNetInfo.NetIOCountersStat[counter.Name].PacketsSent) / diff
		stat.PacketsRecvRate = float64(stat.PacketsRecv-lastNetInfo.NetIOCountersStat[counter.Name].PacketsRecv) / diff
	}

	// 更新上一次的网卡信息
	lastNetIOStatTimeStamp = curTimeStamp
	lastNetInfo = &model.NetInfo{NetIOCountersStat: netMap}

	// 写入influxdb
	writePoints(model.SystemInfo{
		InfoType: consts.NET_INFO_TYPE,
		Data:     model.NetInfo{NetIOCountersStat: netMap},
	})
}
