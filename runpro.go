package main

import (
	"flag"

	"github.com/samverrall/run-pro/config"
)

const (
	defaultRunProConfig = "./runpro-config.json"
)

func main() {
	c := config.New(defaultRunProConfig)

	configOpts, err := c.Read()
	if err != nil {
		panic(err)
	}

	project := flag.String("project", "", "The project to run.")
	flag.Parse()

	if project == nil {
		panic("you must supply a project flag.")
	}
}
