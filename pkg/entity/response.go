package entity

import (
	"log"
	"net/http"
	"time"

	"github.com/olegbespalov/tservice/internal/parser"
)

//Response represent what exactly will be returneds
type Response struct {
	wait       time.Duration
	statusCode int
	body       []byte
	headers    map[string]string
}

//NewResponse creates a new response from the definition
func NewResponse(assetPath string, definition ResponseDefinition) Response {
	wait := time.Nanosecond * 0
	if definition.Slowness != nil && definition.Slowness.Happened() {
		wait, _ = parser.ParseInterval(5*time.Second, definition.Slowness.Duration)
	}

	if definition.Error != nil && definition.Error.Happened() {
		return Response{
			statusCode: definition.Error.StatusCode,
			body:       []byte(`{"error": "yes"}`),
			wait:       wait,
		}
	}

	return Response{
		statusCode: definition.BuildStatusCode(),
		body:       definition.BuildBody(assetPath),
		headers:    definition.BuildHeaders(),
		wait:       wait,
	}
}

//NewDefaultResponse creates a default response
func NewDefaultResponse() Response {
	return Response{
		wait:       0 * time.Nanosecond,
		statusCode: http.StatusOK,
		body:       []byte(`{"deafult": "yes"}`),
	}
}

//Send return response
func (r Response) Send(w http.ResponseWriter) {
	time.Sleep(r.wait)

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
