package cmd

import (
	"fmt"

	"github.com/samverrall/run-pro/config"
	"github.com/samverrall/run-pro/projects"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Attempts to run a project",
	Long:  "Attempts to run a project from the projects config file, using the project name.",
	Run: func(cmd *cobra.Command, args []string) {
		// Read the config to get the location of the projects JSON file.
		c, err := config.Get()
		if err != nil {
			fmt.Printf("list: %v\n", err)
			return
		}

		newProject, err := createNewProjectFromArgs(args)
		switch {
		case err != nil:
			fmt.Println("add: failed to add a project: ", err)
			return
		case newProject == nil:
			fmt.Println("add: got nil newProject")
			return
		}

		c.AddProject(*newProject)
	},
}

func createNewProjectFromArgs(args []string) (*projects.Project, error) {
	// TODO: Handle args mapping to project field, for now just Name is supported in add.
	newProject := projects.Project{
		Name: args[0],
	}
	return &newProject, nil
}
