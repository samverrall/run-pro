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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read the config to get the location of the projects JSON file.
		c, err := config.Get()
		if err != nil {
			return fmt.Errorf("run: %v", err)
		}

		if len(c.Projects) == 0 {
			return fmt.Errorf("no projects exist in your config")
		}

		if len(args) == 0 {
			return fmt.Errorf("no project name supplied")
		}

		project, err := c.Projects.LookupByName(args[0])
		switch {
		case err != nil:
			return fmt.Errorf("failed to lookup project: %v", err)
		case project == nil:
			return fmt.Errorf("run: got nil project")
		default:
			rErr := maybeRunProject(project)
			if rErr != nil {
				return fmt.Errorf("failed to run project: %v", rErr)
			}
		}

		return nil
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
