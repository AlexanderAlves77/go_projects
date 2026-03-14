package cmd

import (
	"fmt"
	"godevops/internal/system"

	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Show memory usage",
	Run: func(cmd *cobra.Command, args []string) {
		usage := system.GetMemoryUsage()
		fmt.Printf("Memory Usage: %.2f%%\n", usage)
	},
}

func init() {
	rootCmd.AddCommand(memoryCmd)
}
