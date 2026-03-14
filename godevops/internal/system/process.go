package system

import (
	"sort"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	Name string
	CPU  float64
}

// GetProcessCount retorna número de processos
func GetProcessCount() int {

	processes, err := process.Processes()

	if err != nil {
		return 0
	}

	return len(processes)
}

// GetTopProcesses retorna os processos que mais usam CPU
func GetTopProcesses(limit int) []ProcessInfo {

	procs, err := process.Processes()
	if err != nil {
		return nil
	}

	var list []ProcessInfo

	for _, p := range procs {

		name, err := p.Name()
		if err != nil {
			continue
		}

		// primeira leitura
		_, err = p.CPUPercent()
		if err != nil {
			continue
		}

		// pequena pausa para calcular o delta
		time.Sleep(50 * time.Millisecond)

		// segunda leitura
		cpu, err := p.CPUPercent()
		if err != nil {
			continue
		}

		list = append(list, ProcessInfo{
			Name: name,
			CPU:  cpu,
		})
	}

	// ordena do maior consumo para o menor
	sort.Slice(list, func(i, j int) bool {
		return list[i].CPU > list[j].CPU
	})

	if len(list) > limit {
		list = list[:limit]
	}

	return list
}
