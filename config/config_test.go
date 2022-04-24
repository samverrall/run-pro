package config

import (
	"reflect"
	"testing"
)

func Test_config_Read(t *testing.T) {
	tests := []struct {
		name    string
		c       *config
		want    *ConfigOptions
		wantErr bool
	}{
		// TODO: Add test cases.
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
