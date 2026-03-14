package cmd

import (
	"fmt"
	"godevops/internal/system"

	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Show network statistics",
	Run: func(cmd *cobra.Command, args []string) {
		sent, recv := system.GetNetworkStats()
		fmt.Printf("Sent: %d bytes\n", sent)
		fmt.Printf("Received: %d bytes\n", recv)
	},
}

func init() {
	rootCmd.AddCommand(networkCmd)
}
