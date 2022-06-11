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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read the config to get the location of the projects JSON file.
		c, err := config.Get()
		if err != nil {
			return fmt.Errorf("list: %v", err)
		}

		if len(c.Projects) == 0 {
			return fmt.Errorf("no projects exist in config file")
		}

		for _, p := range c.Projects {
			fmt.Printf("Name: %s, Entry: %s\n", p.Name, p.EntryFile)
		}

		return nil
	},
}
