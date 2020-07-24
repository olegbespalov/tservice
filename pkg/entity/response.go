package entity

import (
	"log"
	"net/http"
	"time"
)

//Response represent what exactly will be returneds
type Response struct {
	wait       time.Duration
	statusCode int
	body       []byte
	headers    map[string]string
}

//NewResponse creates a new response from the definition
func NewResponse(path string, rule Rule) Response {
	definition := rule.ChoseDefinition()

	return Response{
		statusCode: definition.BuildStatusCode(),
		body:       definition.BuildBody(path),
		headers:    definition.BuildHeaders(),
		wait:       rule.Wait(),
	}
}

//NewDefinitionResponse creates a default response
func NewDefinitionResponse(definition Definition) Response {
	return Response{
		wait:       0 * time.Nanosecond,
		statusCode: definition.BuildStatusCode(),
		body:       []byte(definition.Response),
	}
}

//Send return response
func (r Response) Send(w http.ResponseWriter) {
	time.Sleep(r.wait)

	if len(r.headers) > 0 {
		for k, v := range r.headers {
			w.Header().Set(k, v)
		}
	}

	if r.statusCode > 0 {
		w.WriteHeader(r.statusCode)
	}

	_, err := w.Write(r.body)
	if err != nil {
		log.Printf("ERR during response write: %s", err.Error())
	}
}

//StatusCode http status code of the response
func (r Response) StatusCode() int {
	return r.statusCode
}

//Body of the response
func (r Response) Body() []byte {
	return r.body
}

//Headers of the response
func (r Response) Headers() map[string]string {
	return r.headers
}

//Wait before responding
func (r Response) Wait() time.Duration {
	return r.wait
}
