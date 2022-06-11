package main

import (
	"fmt"

	"github.com/samverrall/run-pro/cmd"
	"github.com/samverrall/run-pro/config"
)

func main() {
	c := config.New(config.DefaultRunProConfig)

	if _, err := c.SetProjects(); err != nil {
		fmt.Println("runpro: failed to read config:", err)
		return
	}

	config.Set(c)

	// Run cobra
	cmd.Execute()
}
