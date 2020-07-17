package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type tservice struct {
}

//NewDefaultHandler return base response
func NewDefaultHandler() http.Handler {
	return &tservice{}
}

func (ts *tservice) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(string(requestDump))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
