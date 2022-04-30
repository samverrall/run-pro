package main

import (
	"fmt"

	"github.com/samverrall/run-pro/cmd"
	"github.com/samverrall/run-pro/config"
)

func main() {
	c := config.New(config.DefaultRunProConfig)

	configOpts, err := c.Read()
	switch {
	case err != nil:
		fmt.Errorf("runpro: failed to read config: %v", err)
	case configOpts == nil:
		fmt.Errorf("runpro: got nil configOpts")
	}

	// Set the config global instance
	config.Set(configOpts)

	// Run cobra
	cmd.Execute()
}
