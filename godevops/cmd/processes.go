package cmd

import (
	"fmt"
	"godevops/internal/system"

	"github.com/spf13/cobra"
)

var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Show running process",
	Run: func(cmd *cobra.Command, args []string) {
		count := system.GetProcessCount()
		fmt.Printf("running processes: %d\n", count)
	},
}

func init() {
	rootCmd.AddCommand(processCmd)
}
