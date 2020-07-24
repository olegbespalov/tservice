package config

import "github.com/olegbespalov/tservice/pkg/entity"

//UseCase config service
type UseCase interface {
	Port() string
	ResponsesPath() string
	Rules() map[string]entity.Rule
	DefaultDefinition() *entity.Definition
	DefaultError() *entity.Definition
}
