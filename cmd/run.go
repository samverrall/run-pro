package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/samverrall/run-pro/config"
	"github.com/samverrall/run-pro/projects"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Attempts to run a project",
	Long:  "Attempts to run a project from the projects config file, using the project name.",
	Run: func(cmd *cobra.Command, args []string) {
		// Read the config to get the location of the projects JSON file.
		c, err := config.Get()
		if err != nil {
			fmt.Printf("run: %v\n", err)
			return
		}

		if len(c.Projects) == 0 {
			fmt.Println("You currently have 0 projects in your config.")
			return
		}

		if len(args) == 0 {
			fmt.Println("You must supply a project name.")
			return
		}

		project, err := c.Projects.LookupByName(args[0])
		switch {
		case err != nil:
			fmt.Println(err)
			return
		case project == nil:
			fmt.Println("run: got nil project")
			return
		default:
			rErr := maybeRunProject(project)
			if rErr != nil {
				fmt.Printf("Failed to run project, error: %v\n", rErr)
				return
			}
		}
	},
}

func maybeRunProject(project *projects.Project) error {
	fmt.Println("Attempting to run project: ", project.Name)
	fmt.Println("Running project in directory: ", project.Dir)
	fmt.Println(project.Args[0], project.Args[1:])

	cmd := exec.Command(project.Args[0], project.Args[1:]...)
	cmd.Dir = project.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
