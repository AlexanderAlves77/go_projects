package system

import (
	"github.com/shirou/gopsutil/v3/disk"
)

// GetDiskUsage retorna o uso do disco principal
func GetDiskUsage() float64 {
	diskstat, err := disk.Usage("C:\\")

	if err != nil {
		return 0
	}

	return diskstat.UsedPercent
}
