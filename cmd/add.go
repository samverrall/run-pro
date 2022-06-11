package cmd

import (
	"errors"
	"fmt"

	"github.com/samverrall/run-pro/config"
	"github.com/samverrall/run-pro/projects"
	"github.com/spf13/cobra"
)

type flags struct {
	Name      string
	Dir       string
	EntryFile string
	Args      []string
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().String("name", "", "The wonderful name of your project")
	addCmd.PersistentFlags().String("dir", "", "The path to your project")
	addCmd.PersistentFlags().String("entry-file", "", "The entry file to start your project")
	addCmd.PersistentFlags().StringSlice("args", []string{}, "The arguments to run your project joined by a ,. `--args=go,run,main.go`")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Attempts to run a project",
	Long:  "Attempts to run a project from the projects config file, using the project name.",
	RunE: func(cmd *cobra.Command, args []string) error {
		flags, err := getCommandFlagValues(cmd)
		if err != nil {
			return err
		}

		// Read the config to get the location of the projects JSON file.
		c, err := config.Get()
		if err != nil {
			return fmt.Errorf("add: %v", err)
		}

		newProject, err := createNewProjectFromArgs(flags)
		switch {
		case err != nil:
			return fmt.Errorf("add: failed to add project: %v", err)
		case newProject == nil:
			return fmt.Errorf("add: got nil newProject")
		}

		if err := c.AddProject(*newProject); err != nil {
			return err
		}

		return nil
	},
}

func createNewProjectFromArgs(f flags) (*projects.Project, error) {
	newProject := projects.Project{
		Name:      f.Name,
		Dir:       f.Dir,
		EntryFile: f.EntryFile,
		Args:      f.Args,
	}
	return &newProject, nil
}

func getCommandFlagValues(cmd *cobra.Command) (flags, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return flags{}, err
	}

	dir, err := cmd.Flags().GetString("dir")
	if err != nil {
		return flags{}, err
	}

	entryFile, err := cmd.Flags().GetString("entry-file")
	if err != nil {
		return flags{}, err
	}

	args, err := cmd.Flags().GetStringSlice("args")
	switch {
	case err != nil:
		return flags{}, err
	case len(args) == 0:
		return flags{}, errors.New("no args supplied to args flag")
	}

	return flags{
		Name:      name,
		Dir:       dir,
		EntryFile: entryFile,
		Args:      args,
	}, nil
}
