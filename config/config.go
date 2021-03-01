package config

import (
	"io/ioutil"
	"os"

	"go.bobheadxi.dev/zapx/zapx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// Config bundles together discord configuration information
type Config struct {
	// the token used by the main bot (NDXBot)
	MainDiscordToken string `yaml:"main_discord_token"`
	// if nil we dont use infura and connect directly to the rpc node below
	InfuraAPIKey      string    `yaml:"infura_api_key"`
	InfuraWSEnabled   bool      `yaml:"infura_ws_enabled"`
	ETHRPCEndpoint    string    `yaml:"eth_rpc_endpoint"`
	MulticallContract string    `yaml:"multicall_contract"`
	Watchers          []Watcher `yaml:"watchers"`
	Database          Database  `yaml:"database"`
	Logger            `yaml:"logger"`
	Indices           []string `yaml:"indices"`
}

// Database provides configuration over our database connection
type Database struct {
	Type           string `yaml:"type"` // sqlite or postgres, if sqlite all other options except DBName are ignored
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Pass           string `yaml:"pass"`
	DBName         string `yaml:"db_name"`
	DBPath         string `yaml:"db_path"`
	SSLModeDisable bool   `yaml:"ssl_mode_disable"`
}

// Watcher is used to start a process that watches the price of a token
// and posts its value as a name
type Watcher struct {
	DiscordToken string `yaml:"discord_token"`
	Currency     string `yaml:"currency"`
}

// Logger provides configuration over zap logger
type Logger struct {
	Path     string                 `yaml:"path"`
	Debug    bool                   `yaml:"debug"`
	Dev      bool                   `yaml:"dev"`
	FileOnly bool                   `yaml:"file_only"`
	Fields   map[string]interface{} `yaml:"fields"`
}

var (
	// ExampleConfig is primarily used to provide a template for generating the config file
	ExampleConfig = &Config{
		MainDiscordToken:  "CHANGEME-MAIN",
		InfuraAPIKey:      "INFURA-KEY",
		InfuraWSEnabled:   false,
		ETHRPCEndpoint:    "http://localhost:8545",
		MulticallContract: "0xFB6FdE35dDD2cC295908b1E5EfAF36Effb5C04dD", // NOTE: this is the multicall from the circuit breaker package and this needs to be updated
		Watchers: []Watcher{
			{DiscordToken: "CHANGEME-TOKEN", Currency: "CHANGEME-CURRENCY"},
		},
		Database: Database{
			Type:           "sqlite",
			Host:           "localhost",
			Port:           "5432",
			User:           "user",
			Pass:           "pass",
			DBName:         "indexed",
			DBPath:         "/changeme",
			SSLModeDisable: false,
		},
		Indices: []string{"defi5", "cc10", "orcl5", "degen10"},
		Logger: Logger{
			Path:  "gondx.log",
			Debug: true,
			Dev:   true,
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

// LoggerFromConfig returns a logger from our config
func LoggerFromConfig(cfg *Config) (*zap.Logger, error) {
	var opts []zapx.Option
	if cfg.Logger.Debug {
		opts = append(opts, zapx.WithDebug(true))
	}
	if cfg.Logger.FileOnly {
		opts = append(opts, zapx.OnlyToFile())
	}
	if cfg.Logger.Fields != nil {
		opts = append(opts, zapx.WithFields(cfg.Logger.Fields))
	}
	return zapx.New(
		cfg.Logger.Path,
		cfg.Logger.Dev,
		opts...,
	)
}
