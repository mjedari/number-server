package configs

import "testing"

func TestGetConfigs(t *testing.T) {
	setConfigPath(".")
	config := GetConfigs()

	if config == nil {
		t.Error("failed to load config")
	}
}
