package entity

// Config for the service
type Config struct {
	ResponseRules map[string]ResponseRules `yaml:"responses"`
}
