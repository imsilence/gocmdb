package entity

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Resource struct {
	Load           *load.AvgStat                  `json:"load"`
	CPUPrecent     float64                        `json:"cpu_precent"`
	MemPrecent     float64                        `json:"mem_precent"`
	DiskPrecent    map[string]float64             `json:"disk_precent"`
	DiskIOCounters map[string]disk.IOCountersStat `json:"disk_io_counters"`
	NetIOCounters  []net.IOCountersStat           `json:"net_io_counters"`
}

func NewResource(uuid string) Resource {
	loadInfo, _ := load.Avg()
	cpuPrecents, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	partitionInfo, _ := disk.Partitions(true)

	diskPrecents := map[string]float64{}

	for _, partition := range partitionInfo {
		usageInfo, _ := disk.Usage(partition.Device)
		diskPrecents[usageInfo.Path] = usageInfo.UsedPercent
	}

	diskIOCounters, _ := disk.IOCounters()
	netIOCounters, _ := net.IOCounters(false)

	return Resource{
		Load:           loadInfo,
		CPUPrecent:     cpuPrecents[0],
		MemPrecent:     memInfo.UsedPercent,
		DiskPrecent:    diskPrecents,
		DiskIOCounters: diskIOCounters,
		NetIOCounters:  netIOCounters,
	}
}
