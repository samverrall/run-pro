package projects

type Project struct {
	Name      string   `json:"name"`
	Dir       string   `json:"dir"`
	EntryFile string   `json:"entryFile"`
	Args      []string `json:"args"`
}

type ProjectsIn []Project

func Lookup(name string) {

}
