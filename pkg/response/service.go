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
	for _, responseDefinition := range s.cfg.Config().ResponseDefinitions {
		if responseDefinition.Fit(r.Method, r.RequestURI) {
			return entity.NewResponse(s.cfg.AssetPath(), responseDefinition)
		}
	}

	return entity.NewDefaultResponse()
}
