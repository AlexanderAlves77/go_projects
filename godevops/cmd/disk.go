package cmd

import (
	"fmt"
	"godevops/internal/system"

	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Show disk usage",
	Run: func(cmd *cobra.Command, args []string) {
		usage := system.GetDiskUsage()
		fmt.Printf("Disk Usage: %.2f%%\n", usage)
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)
}
