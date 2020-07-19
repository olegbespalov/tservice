package response

import (
	"net/http"

	"github.com/olegbespalov/tservice/pkg/entity"
)

//UseCase response service interface
type UseCase interface {
	BestResponse(r *http.Request) entity.Response
}
