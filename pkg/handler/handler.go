package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/olegbespalov/tservice/pkg/config"
)

type service struct {
	cfg config.UseCase
}

//NewDefaultHandler return base response
func NewDefaultHandler(cfg config.UseCase) http.Handler {
	return &service{
		cfg: cfg,
	}
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(string(requestDump))

	for _, res := range s.cfg.Responses() {
		if res.Fit(r.Method, r.RequestURI) {
			res.Send(w)

			return
		}
	}

	_, err = w.Write([]byte(`{"status": "ok"}`))
	if err != nil {
		log.Println("ERR: " + err.Error())
	}
}
