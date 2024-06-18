package influx

import (
	"LogCollector/logic/consts"
	"LogCollector/logic/model"
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"time"
)

var (
	CLIENT             influxdb2.Client
	INFLUXDB_WRITE_API api.WriteAPIBlocking
)

// InfluxDbConn
// @Description: 初始化influxdb连接
func InfluxDbConn(org string, buket string) {
	token := "xCs_T4YxAzOWC2IxyJcTp_ExGynRXQgDFqjXy8opWPULIhCKJYL8SMu9HdvG78Kkta970pLflgivRZECsBb_Ew=="
	// Store the URL of your InfluxDB instance
	url := "http://localhost:8086"
	// Create new cli with default option for server url authenticate by token
	CLIENT = influxdb2.NewClient(url, token)
	// User blocking write client for writes to desired bucket
	INFLUXDB_WRITE_API = CLIENT.WriteAPIBlocking(org, buket)
}

func writePoints(info model.SystemInfo) {
	// 根据传入的类型，转换数据
	switch info.InfoType {
	case consts.CPU_INFO_TYPE:
		writePointCpu(info)
		break
	case consts.MEM_INFO_TYPE:
		writePointMem(info)
		break
	case consts.DISK_INFO_TYPE:
		writePointDisk(info)
		break
	case consts.NET_INFO_TYPE:
		writePointNet(info)
		break
	}
}

func writePointCpu(info model.SystemInfo) {
	measurement := "cpu"
	tags := map[string]string{"cpu": "cpu0"}
	data := info.Data.(model.CpuInfo)
	fields := map[string]interface{}{
		"cpu_percent": data.CpuPercent,
	}

	// Create point using full params constructor
	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	// Write point immediately
	INFLUXDB_WRITE_API.WritePoint(context.Background(), p)
}

func writePointMem(info model.SystemInfo) {
	measurement := "memory"
	tags := map[string]string{"mem": "memory"}
	data := info.Data.(model.MemInfo)
	fields := map[string]interface{}{
		"total":        int64(data.Total),
		"available":    int64(data.Available),
		"used":         int64(data.Used),
		"used_percent": data.UsedPercent,
		"buffers":      int64(data.Buffers),
		"cached":       int64(data.Cached),
	}

	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	INFLUXDB_WRITE_API.WritePoint(context.Background(), p)
}

func writePointDisk(info model.SystemInfo) {
	measurement := "disk"
	data := info.Data.(model.DiskInfo)
	for key, value := range data.PartitionUsageStat {
		tags := map[string]string{"disk": key}
		fields := map[string]interface{}{
			"total":               value.Total,
			"free":                value.Free,
			"used":                value.Used,
			"used_percent":        value.UsedPercent,
			"inodes_total":        value.InodesTotal,
			"inodes_used":         value.InodesUsed,
			"inodes_free":         value.InodesFree,
			"Inodes_used_percent": value.InodesUsedPercent,
		}
		// Create point using full params constructor
		p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
		// Write point immediately
		INFLUXDB_WRITE_API.WritePoint(context.Background(), p)
	}
}

func writePointNet(info model.SystemInfo) {
	measurement := "net"
	data := info.Data.(model.NetInfo)
	for key, value := range data.NetIOCountersStat {
		tags := map[string]string{"name": key}
		fields := map[string]interface{}{
			"BytesSentRate":   value.BytesSentRate,
			"BytesRecveRate":  value.BytesRecveRate,
			"PacketsSentRate": value.PacketsSentRate,
			"PacketsRecvRate": value.PacketsRecvRate,
		}
		// Create point using full params constructor
		p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
		// Write point immediately
		INFLUXDB_WRITE_API.WritePoint(context.Background(), p)
	}
	return
}
