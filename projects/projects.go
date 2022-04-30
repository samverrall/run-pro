package projects

// Project defines the fields available on the JSON project object.
type Project struct {
	Name      string   `json:"name"`
	EntryFile string   `json:"entryFile"`
	Args      []string `json:"args"`
}

type ProjectsIn []Project
