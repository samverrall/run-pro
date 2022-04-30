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
		fmt.Println("runpro: failed to read config:", err)
		return
	case configOpts == nil:
		fmt.Println("runpro: got nil configOpts")
		return
	}

	// Set the config global instance
	config.Set(configOpts)

	// Run cobra
	cmd.Execute()
}
