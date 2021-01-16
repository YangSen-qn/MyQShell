package common

var (
	OutputFormatDefault = "default"
	OutputFormatJSON    = "json"
	OutputFormatXML     = "xml"
	OutputFormatYAML    = "yaml"
)

type Config struct {
	OutputFormatValue string
}

