package system

import (
	"github.com/shirou/gopsutil/v3/mem"
)

// GetMemoryUsage retorna o percentual de uso da memória RAM
func GetMemoryUsage() float64 {

	vmStat, err := mem.VirtualMemory()

	if err != nil {
		return 0
	}

	return vmStat.UsedPercent
}
