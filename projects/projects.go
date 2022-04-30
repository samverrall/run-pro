package projects

import (
	"fmt"
	"strings"
)

// Project defines the fields available on the JSON project object.
type Project struct {
	Name      string   `json:"name"`
	Dir       string   `json:"dir"`
	EntryFile string   `json:"entryFile"`
	Args      []string `json:"args"`
}

type ProjectsIn []Project

func (pi ProjectsIn) LookupByName(name string) (*Project, error) {
	for _, p := range pi {
		if strings.EqualFold(p.Name, name) {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("project %s does not exist", name)
}
