package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godevops",
	Short: "DevOps CLI built with Go",
	Long:  "A command line tool to inspect system resources.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a command like cpu or memory")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
