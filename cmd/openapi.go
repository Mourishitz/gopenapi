package main

type OpenAPI struct {
	OpenAPI      string                   `yaml:"openapi"`
	Info         map[string]interface{}   `yaml:"info,omitempty"`
	Security     []map[string]interface{} `yaml:"security,omitempty"`
	Tags         []map[string]interface{} `yaml:"tags,omitempty"`
	Paths        map[string]interface{}   `yaml:"paths,omitempty"`
	ExternalDocs map[string]interface{}   `yaml:"externalDocs,omitempty"`
	Servers      []map[string]interface{} `yaml:"servers,omitempty"`
	Components   map[string]interface{}   `yaml:"components,omitempty"`
	Definitions  map[string]interface{}   `yaml:"definitions,omitempty"`
}
