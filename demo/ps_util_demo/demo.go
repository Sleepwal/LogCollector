package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func main() {
	getCpuInfo()
	fmt.Println("------------------")
	go getCpuLoad()
	fmt.Println("------------------")
	getMemInfo()
	fmt.Println("------------------")
	getHostInfo()
	fmt.Println("------------------")
	getDiskInfo()
	fmt.Println("------------------")
	getNetInfo()
	fmt.Println("------------------")
}

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
		time.Sleep(5 * time.Second)
	}
}

// cpu 负载
func getCpuLoad() {
	avg, err := load.Avg()
	if err != nil {
		fmt.Printf("get cpu load failed, err:%v", err)
	}
	fmt.Printf("cpu load:%v\n", avg)
}

// 内存信息
func getMemInfo() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v", err)
	}
	fmt.Printf("mem info:%v\n", memInfo)
}

// Host info
func getHostInfo() {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Printf("get host info failed, err:%v", err)
	}
	fmt.Printf("host info:%v\n", hostInfo)
}

// 磁盘信息
func getDiskInfo() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get disk info failed, err:%v", err)
	}
	for _, part := range partitions {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			fmt.Printf("get disk info failed, err:%v", err)
			return
		}
		fmt.Printf("disk info:%v\n", usage)
	}
}

// 网络信息
func getNetInfo() {
	counters, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net info failed, err:%v", err)
		return
	}
	for _, c := range counters {
		fmt.Printf("net info:%v\n", c)
	}
}
