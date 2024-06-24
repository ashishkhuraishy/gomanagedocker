package config

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knadh/koanf/v2"
)

func TestReadConfig(t *testing.T) {

	tests := []struct {
		UserConfig string
		Want       map[string]any
	}{
		{
			UserConfig: "",
			Want:       map[string]any{"config.Polling-Time": 500, "config.Tab-Order": []any{"images", "containers", "volumes"}},
		},
		{
			UserConfig: `config:
  Polling-Time: 100`,
			Want: map[string]any{"config.Polling-Time": 100, "config.Tab-Order": []any{"images", "containers", "volumes"}},
		},
		{
			UserConfig: `config:
  Polling-Time: 200
  Tab-Order: [containers, volumes]`,
			Want: map[string]any{"config.Polling-Time": 200, "config.Tab-Order": []any{"containers", "volumes"}},
		},
	}

	for id, test := range tests {
		tempFile, _ := os.CreateTemp("", "")
		tempFile.WriteString(test.UserConfig)
		defer os.Remove(tempFile.Name())

		got := koanf.New(".")
		filePath := tempFile.Name()
		ReadConfig(got, filePath)

		if !cmp.Equal(got.All(), test.Want) {
			t.Errorf("Fail %d: %s", id, cmp.Diff(got.All(), test.Want))
		}

	}

}
