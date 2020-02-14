package entity

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Register struct {
	UUID     string    `json:"uuid"`
	Hostname string    `json:"hostname"`
	IP       string    `json:"ip"`
	OS       string    `json:"os"`
	Arch     string    `json:"arch"`
	Ram      int64     `json:"ram"`
	CPU      int       `json:"cpu"`
	Disk     string    `json:"disk"`
	Boottime time.Time `json:"boottime"`
	Time     time.Time `json:"time"`
}

func NewRegister(uuid string) Register {
	hostInfo, _ := host.Info()

	ips := []string{}

	interfaceStat, _ := net.Interfaces()
	for _, intf := range interfaceStat {
		for _, addr := range intf.Addrs {
			if strings.Index(addr.Addr, ":") >= 0 {
				continue
			}
			if strings.Index(addr.Addr, "127.") == 0 {
				continue
			}
			nodes := strings.SplitN(addr.Addr, "/", 2)
			ips = append(ips, nodes[0])
		}
	}

	cores := 0
	cpuInfo, _ := cpu.Info()

	for _, cpu := range cpuInfo {
		cores += int(cpu.Cores)
	}

	memInfo, _ := mem.VirtualMemory()

	boottime, _ := host.BootTime()
	partitions := map[string]int64{}

	partitionInfo, _ := disk.Partitions(true)

	for _, partition := range partitionInfo {
		usageInfo, _ := disk.Usage(partition.Device)
		partitions[usageInfo.Path] = int64(usageInfo.Total)
	}

	ipJson, _ := json.Marshal(ips)
	diskJson, _ := json.Marshal(partitions)
	return Register{
		UUID:     uuid,
		Hostname: hostInfo.Hostname,
		IP:       string(ipJson),
		OS:       hostInfo.OS,
		CPU:      cores,
		Ram:      int64(memInfo.Total),
		Disk:     string(diskJson),
		Boottime: time.Unix(int64(boottime), 0),
		Time:     time.Now(),
	}
}
