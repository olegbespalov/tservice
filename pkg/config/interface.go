package config

import "github.com/olegbespalov/tservice/pkg/entity"

//UseCase config service
type UseCase interface {
	Config() entity.Config
	Responses() map[string]entity.Response
}
