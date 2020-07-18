package entity

import (
	"log"
	"net/http"
)

var defaultBody = []byte(`{"status": "ok"}`)

//Response service response
type Response struct {
	Method     string
	Path       string
	StatusCode int `yaml:"status_code"`

	Body []byte

	Response     string
	ResponseFile string `yaml:"response_file"`

	Headers  map[string]string
	Error    Error
	Slowness Slowness
}

//Error define if response will be error
type Error struct {
	Chance     int
	Time       string
	StatusCode int `yaml:"status_code"`
}

//Slowness define if response will be slow
type Slowness struct {
	Chance int
	Time   string
}

// Fit check if response can be used for the request
func (r Response) Fit(method, path string) bool {
	return r.Path == path && (len(r.Method) == 0 || r.Method == method)
}

//Send return response
func (r Response) Send(w http.ResponseWriter) {
	if r.StatusCode > 0 {
		w.WriteHeader(r.StatusCode)
	}

	_, err := w.Write(r.getBody())
	if err != nil {
		log.Printf("ERR during response write: %s", err.Error())
	}
}

func (r Response) getBody() []byte {
	if len(r.Body) > 0 {
		return r.Body
	}

	return defaultBody
}
