package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/samverrall/run-pro/projects"
)

const (
	DefaultRunProConfig = "./runpro-config.json"
)

type ConfigOptions struct {
	Projects projects.ProjectsIn `json:"projects"`
}

type config struct {
	file string
}

func New(file string) *config {
	return &config{
		file: file,
	}
}

func (c *config) Read() (*ConfigOptions, error) {
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
		return nil, errors.New("config file not in proper json formats")
	}

	return &configOpts, nil
}
