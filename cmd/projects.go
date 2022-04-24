package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(projectCmd)
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Print the version number of run-rpo",
	Long:  `All software has versions. This is runpro's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RunPro version v0.0.1", args)
	},
}
