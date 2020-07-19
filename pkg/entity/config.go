package entity

// Config for the service
type Config struct {
	ResponseDefinitions map[string]ResponseDefinition `yaml:"responses"`
}
