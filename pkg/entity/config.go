package entity

// Config for the service
type Config struct {
	Default map[string]Definition `yaml:"default_definitions"`
	Rules   map[string]Rule       `yaml:"responses"`
}
