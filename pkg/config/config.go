package config

import (
	"os"

	"github.com/gleich/statuser/v2"
	"gopkg.in/yaml.v3"
)

var path = "/.config/nuke/"

// Config for nuke
type Conf struct {
	IgnoreUpdates      bool     `yaml:"ignoredUpdates"`
	IgnoredApps        []string `yaml:"ignoredApps"`
	IgnoredRunningFrom bool     `yaml:"ignoredRunningFrom"`
}

// Check if the config exists
func Exists() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get user home path", err, 1)
	}
	path = home + path

	yml, err := os.Stat(path + "config.yml")
	if !os.IsNotExist(err) {
		if !yml.IsDir() {
			path = path + "config.yml"
			return true
		}
	}
	yaml, err := os.Stat(path + "config.yaml")
	if !os.IsNotExist(err) {
		if !yaml.IsDir() {
			path = path + "config.yaml"
			return true
		}
	}
	return false
}

// Read from the config file
func Read(conf *Conf) *Conf {
	contents, err := os.ReadFile(path)
	if err != nil {
		statuser.Error("Failed to read from config in\n\t"+path, err, 1)
	}

	err = yaml.Unmarshal(contents, &conf)
	if err != nil {
		statuser.Error("Failed to read the config", err, 1)
	}
	return conf
}
