package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/samverrall/run-pro/config"
)

const (
	defaultRunProConfig = "./runpro-config.json"
)

type flags struct {
	project string
}

func parseFlags() (*flags, error) {
	project := flag.String("project", "", "The project to run.")
	flag.Parse()

	if project == nil {
		return nil, errors.New("no project flag supplied")
	}

	return &flags{
		project: *project,
	}, nil
}

func main() {
	c := config.New(defaultRunProConfig)

	configOpts, err := c.Read()
	switch {
	case err != nil:
		fmt.Errorf("runpro: failed to read config: %v", err)
	case configOpts == nil:
		fmt.Errorf("runpro: got nil configOpts")
	}

	flags, err := parseFlags()
	if err != nil {
		fmt.Errorf("runpro: failed to parse flags: %v", err)
	}

	// TODO: Read flags and implement project runner, and other commmands.
	fmt.Println(flags)
}
