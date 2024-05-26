package cpu

import (
    "time"

    "github.com/shirou/gopsutil/cpu"
)

type cpuStatsStruct struct {
    LoadAvg float64
}

var cpuStats = cpuStatsStruct{
    LoadAvg: 0,
}

func GetCPUStats() cpuStatsStruct {
    return cpuStats 
}

func UpdateCPUStats()  {
    loadAvg, _ := cpu.Percent(3 * time.Second, false)
    cpuStats.LoadAvg = loadAvg[0]
}
