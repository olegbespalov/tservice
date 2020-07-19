package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/olegbespalov/tservice/pkg/response"
)

type service struct {
	responses response.UseCase
}

//NewDefaultHandler return base response
func NewDefaultHandler(responses response.UseCase) http.Handler {
	return &service{
		responses: responses,
	}
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(string(requestDump))

	s.responses.BestResponse(r).Send(w)
}
