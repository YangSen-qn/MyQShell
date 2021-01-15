package common

import (
	"github.com/spf13/cobra"
)

var (
	OutputFormatDefault = "default"
	OutputFormatJSON    = "json"
	OutputFormatXML     = "xml"
	OutputFormatYAML    = "yaml"
)

type Config struct {
	OutputFormatValue string
}

func ConfigParam(command *cobra.Command, config *Config) {
	command.Flags().StringVarP(&config.OutputFormatValue, "outputFormat", "", "", "")
}
