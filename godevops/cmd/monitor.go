package cmd

import (
	"fmt"
	"godevops/internal/system"
	"time"

	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Real-time system monitor",
	Run: func(cmd *cobra.Command, args []string) {

		for {
			cpu := system.GetCPUUsage()
			mem := system.GetMemoryUsage()
			disk := system.GetDiskUsage()
			proc := system.GetProcessCount()
			up, down := system.GetNetworkSpeed()

			upMB := system.BytesToMB(up)
			downMB := system.BytesToMB(down)

			fmt.Print("\033[H\033[2J")

			fmt.Printf("\rCPU: %.2f%% | RAM: %.2f%% | DISK: %.2f%% | PROC: %d | NET ↓ %.2fMB/s ↑ %.2fMB/s",
				cpu, mem, disk, proc, downMB, upMB)

			time.Sleep(1 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
