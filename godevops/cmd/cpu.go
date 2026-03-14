package cmd

import (
	"fmt"
	"godevops/internal/system"

	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Show CPU usage",
	Run: func(cmd *cobra.Command, args []string) {
		usage := system.GetCPUUsage()
		fmt.Printf("CPU Usage: %.2f%%\n", usage)
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
}
