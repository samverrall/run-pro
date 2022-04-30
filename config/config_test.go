package config

import (
	"reflect"
	"testing"

	"github.com/samverrall/run-pro/projects"
)

func Test_config_Read(t *testing.T) {
	const (
		successfulConfigFile = "./testdata/success.json"
		invalidConfigFile    = "./testdata/invalid.json"
	)

	tests := []struct {
		name    string
		c       *Config
		want    *ConfigOptions
		wantErr bool
	}{
		{
			name: "Successful config read",
			c:    New(successfulConfigFile),
			want: &ConfigOptions{
				Projects: projects.ProjectsIn{
					{
						Name:      "success",
						EntryFile: "success",
						Args:      []string{"1", "2", "3"},
					},
				},
			},
		},
		{
			name:    "Returns invalid config error",
			c:       New(invalidConfigFile),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("config.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("config.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
