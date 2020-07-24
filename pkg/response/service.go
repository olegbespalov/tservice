package response

import (
	"net/http"

	"github.com/olegbespalov/tservice/pkg/config"
	"github.com/olegbespalov/tservice/pkg/entity"
)

type service struct {
	cfg config.UseCase
}

//NewService create a new response service
func NewService(cfg config.UseCase) UseCase {
	return service{
		cfg: cfg,
	}
}

//BestResponse find the best response
func (s service) BestResponse(r *http.Request) entity.Response {
	for _, responseRule := range s.cfg.Rules() {
		if responseRule.Fit(r.Method, r.RequestURI) {
			return entity.NewResponse(s.cfg.ResponsesPath(), responseRule)
		}

	}

	return entity.NewDefinitionResponse(s.DefaultDefinition())
}

//DefaultDefinition return default definition
func (s *service) DefaultDefinition() entity.Definition {
	if v := s.cfg.DefaultDefinition(); v != nil {
		return *v
	}

	return entity.Definition{
		StatusCode: http.StatusOK,
		Response:   "{\"default\":\"default\"}",
		Headers: []string{
			"Content-Type:application/json",
		},
	}
}

//DefaultError return default error if defined
func (s *service) DefaultError() entity.Definition {
	if v := s.cfg.DefaultError(); v != nil {
		return *v
	}

	return entity.Definition{
		StatusCode: http.StatusInternalServerError,
		Response:   "{\"error\":\"internal server error\"}",
		Headers: []string{
			"Content-Type:application/json",
		},
	}
}
