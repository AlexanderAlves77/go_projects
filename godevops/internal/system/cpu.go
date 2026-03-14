package system

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

// GetCPUUsage retorna o percentual de uso da CPU
func GetCPUUsage() float64 {
	percentages, err := cpu.Percent(0, false)

	if err != nil {
		return 0
	}

	return percentages[0]
}
