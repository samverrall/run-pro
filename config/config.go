package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/samverrall/run-pro/projects"
)

const (
	DefaultRunProConfig = "./runpro-config.json"
)

var ConfigInst *ConfigOptions

func Set(c *ConfigOptions) {
	ConfigInst = c
}

func Get() (*ConfigOptions, error) {
	if ConfigInst == nil {
		return nil, errors.New("cannot get config instance as it is nil")
	}
	return ConfigInst, nil
}

type ConfigOptions struct {
	Projects projects.ProjectsIn `json:"projects"`
}

type Config struct {
	file     string
	Projects projects.ProjectsIn `json:"projects"`
}

func New(file string) *Config {
	return &Config{
		file: file,
	}
}

func (c *Config) Read() (*ConfigOptions, error) {
	bytes, err := os.ReadFile(c.file)
	if err != nil {
		return nil, errors.New("failed to read config file")
	}

	if len(bytes) == 0 {
		return nil, errors.New("empty config file supplied")
	}

	var configOpts ConfigOptions
	jErr := json.Unmarshal(bytes, &configOpts)
	if jErr != nil {
		return nil, errors.New("config file not in proper json format")
	}

	return &configOpts, nil
}

// TODO: Move all ConfigOptions types and methods to a configoptions.go file.
func (c *ConfigOptions) AddProject(project projects.Project) (*ConfigOptions, error) {
	// TODO: JSON encode config options, and write to the file.
	fmt.Println("to add", project)
	return nil, nil
}
