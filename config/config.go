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

var ConfigInst *Config

type Configer interface {
	AddProject(project projects.Project) error
	ToJSON(c *Config) ([]byte, error)
	Write(jsonBytes []byte) error
}

type GetterSetter interface {
	Set(c *Config)
	Get() (*Config, error)
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

func (c *Config) SetProjects() (*Config, error) {
	bytes, err := os.ReadFile(c.file)
	if err != nil {
		return nil, errors.New("failed to read config file")
	}

	if len(bytes) == 0 {
		return nil, errors.New("empty config file supplied")
	}

	var configOpts Config
	jErr := json.Unmarshal(bytes, &configOpts)
	if jErr != nil {
		return nil, errors.New("config file not in proper json format")
	}

	c.Projects = configOpts.Projects

	return c, nil
}

func (c *Config) AddProject(project projects.Project) error {
	newProjectName := project.Name
	existingProject, err := c.Projects.LookupByName(newProjectName)
	switch {
	case err != nil:
		return fmt.Errorf("failed to determind if project already exists: %v", err)
	case existingProject != nil:
		return fmt.Errorf("cannot add project with name %s because a project already exists called %s", newProjectName, newProjectName)
	}

	c.Projects = append(c.Projects, project)

	jsonBytes, err := c.ToJSON()
	if err != nil {
		return err
	}

	if err := c.OverwriteFile(jsonBytes); err != nil {
		return err
	}

	return nil
}

func (c *Config) ToJSON() ([]byte, error) {
	configBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return configBytes, nil
}

func (c *Config) OverwriteFile(jsonBytes []byte) error {
	file, err := os.OpenFile(c.file, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(string(jsonBytes)); err != nil {
		return err
	}

	return nil
}

func Set(c *Config) {
	ConfigInst = c
}

func Get() (*Config, error) {
	if ConfigInst == nil {
		return nil, errors.New("cannot get config instance as it is nil")
	}
	return ConfigInst, nil
}
