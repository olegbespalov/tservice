package config

import "github.com/olegbespalov/tservice/pkg/entity"

//UseCase config service
type UseCase interface {
	Port() string
	ResponsesPath() string
	Config() entity.Config
	ResponseRules() map[string]entity.ResponseRules
}
