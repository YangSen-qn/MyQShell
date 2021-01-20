package common

import "qshell/iqshell/config"

var (
	OutputFormatDefault = "default"
	OutputFormatJSON    = "json"
	OutputFormatXML     = "xml"
	OutputFormatYAML    = "yaml"
)

type Config struct {
	LoadConfig *config.LoadConfig

	CustomConfigPath  string
	OutputFormatValue string
}

func NewConfig() *Config {
	return &Config{
		LoadConfig: &config.LoadConfig{
			RootDir: "",
		},
		OutputFormatValue: OutputFormatDefault,
	}
}
