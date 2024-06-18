package model

import (
	"github.com/shirou/gopsutil/disk"
)

type SystemInfo struct {
	InfoType string      `json:"info_type"`
	Data     interface{} `json:"data"`
}

type CpuInfo struct {
	CpuPercent float64 `json:"cpu_percent"`
}

type MemInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Buffers     uint64  `json:"buffers"`
	Cached      uint64  `json:"cached"`
}

type DiskInfo struct {
	PartitionUsageStat map[string]*disk.UsageStat
}

type IOStat struct {
	BytesSent       uint64  `json:"bytesSent"`   // number of bytes sent
	BytesRecv       uint64  `json:"bytesRecv"`   // number of bytes received
	PacketsSent     uint64  `json:"packetsSent"` // number of packets sent
	PacketsRecv     uint64  `json:"packetsRecv"` // number of packets received
	BytesSentRate   float64 `json:"bytes_sent_rate"`
	BytesRecveRate  float64 `json:"bytes_recve_rate"`
	PacketsSentRate float64 `json:"packets_sent_rate"`
	PacketsRecvRate float64 `json:"packets_recv_rate"`
}

type NetInfo struct {
	NetIOCountersStat map[string]*IOStat
}
