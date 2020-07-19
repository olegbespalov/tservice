package entity

import (
	"log"
	"net/http"
)

//Response represent what exactly will be returneds
type Response struct {
	statusCode int
	body       []byte
	headers    map[string]string
}

//NewResponse creates a new response from the definition
func NewResponse(assetPath string, definition ResponseDefinition) Response {
	return Response{
		statusCode: definition.BuildStatusCode(),
		body:       definition.BuildBody(assetPath),
		headers:    definition.BuildHeaders(),
	}
}

//NewDefaultResponse creates a default response
func NewDefaultResponse() Response {
	return Response{
		statusCode: http.StatusOK,
		body:       []byte(`{"deafult": "yes"}`),
	}
}

//Send return response
func (r Response) Send(w http.ResponseWriter) {
	if r.statusCode > 0 {
		w.WriteHeader(r.statusCode)
	}

	_, err := w.Write(r.body)
	if err != nil {
		log.Printf("ERR during response write: %s", err.Error())
	}
}
