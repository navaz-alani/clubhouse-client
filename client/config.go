package client

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"gopkg.in/yaml.v2"
)

type ClientConfig struct {
	BaseURL  string `yaml:"base_url"`
	ApiToken string `yaml:"api_token"`
}

// LoadConfig loads the client configuration from the ~/.config/clubhouse
// directory, from a file called client.conf. The config is specified in YAML
// format, providing the base_url and api_token keys for the client.
func LoadConfig() (*ClientConfig, error) {
	usr, err := user.Current()
	configFile := fmt.Sprintf("%s/.config/clubhouse/client.conf", usr.HomeDir)

	config := new(ClientConfig)
	yamlConfig, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlConfig, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
