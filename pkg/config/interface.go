package config

import "github.com/olegbespalov/tservice/pkg/entity"

//UseCase config service
type UseCase interface {
	AssetPath() string
	Config() entity.Config
	ResponseDefinition() map[string]entity.ResponseDefinition
}
