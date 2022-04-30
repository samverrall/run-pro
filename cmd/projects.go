package cmd

import (
	"fmt"

	"github.com/samverrall/run-pro/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a list of projects in the config.",
	Run: func(cmd *cobra.Command, args []string) {
		// Read the config to get the location of the projects JSON file.
		c, err := config.Get()
		if err != nil {
			fmt.Printf("list: %v\n", err)
			return
		}

		if len(c.Projects) == 0 {
			fmt.Println("You currently have 0 projects in your config.")
			return
		}

		for _, p := range c.Projects {
			fmt.Printf("Name: %s, Entry: %s", p.Name, p.EntryFile)
		}
	},
}
