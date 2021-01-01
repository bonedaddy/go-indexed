package discord

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config bundles together discord configuration information
type Config struct {
	// the token used by the main bot (NDXBot)
	MainDiscordToken string `yaml:"main_discord_token"`
	// if nil we dont use infura and connect directly to the rpc node below
	InfuraAPIKey    string    `yaml:"infura_api_key"`
	InfuraWSEnabled bool      `yaml:"infura_ws_enabled"`
	ETHRPCEndpoint  string    `yaml:"eth_rpc_endpoint"`
	Watchers        []Watcher `yaml:"watchers"`
}

// Watcher is used to start a process that watches the price of a token
// and posts its value as a name
type Watcher struct {
	DiscordToken string `yaml:"discord_token"`
	Currency     string `yaml:"currency"`
}

var (
	// ExampleConfig is primarily used to provide a template for generating the config file
	ExampleConfig = &Config{
		MainDiscordToken: "CHANGEME-MAIN",
		InfuraAPIKey:     "INFURA-KEY",
		InfuraWSEnabled:  false,
		ETHRPCEndpoint:   "http://localhost:8545",
		Watchers: []Watcher{
			{DiscordToken: "CHANGEME-TOKEN", Currency: "CHANGEME-CURRENCY"},
		},
	}
)

// NewConfig generates a new config and stores at path
func NewConfig(path string) error {
	data, err := yaml.Marshal(ExampleConfig)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, os.ModePerm)
}

// LoadConfig loads the configuration
func LoadConfig(path string) (*Config, error) {
	r, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(r, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
