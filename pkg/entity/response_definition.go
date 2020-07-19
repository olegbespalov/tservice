package entity

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

//ResponseDefinition service response
type ResponseDefinition struct {
	Method string
	Path   string

	StatusCode   int `yaml:"status_code"`
	Response     string
	ResponseFile string `yaml:"response_file"`
	Body         []byte
	Headers      map[string]string

	Error    *Error
	Slowness *Slowness
}

//Error define if response will be error
type Error struct {
	Chance     int
	StatusCode int `yaml:"status_code"`
}

//Happened check if error happened
func (e Error) Happened() bool {
	return rand.Intn(100) <= e.Chance
}

//Slowness define if response will be slow
type Slowness struct {
	Chance   int
	Duration string
}

//Happened check if slowness happened
func (s Slowness) Happened() bool {
	return rand.Intn(100) <= s.Chance
}

// Fit check if response can be used for the request
func (r ResponseDefinition) Fit(method, path string) bool {
	return r.Path == path && (len(r.Method) == 0 || r.Method == method)
}

// BuildBody define what will be in the body
func (r *ResponseDefinition) BuildBody(assetPath string) []byte {
	if len(r.Response) > 0 {
		return []byte(r.Response)
	}

	if len(r.ResponseFile) == 0 {
		return []byte{}
	}

	data, err := ioutil.ReadFile(filepath.Clean(assetPath + string(os.PathSeparator) + r.ResponseFile))
	if err != nil {
		log.Printf("can't find a response file %s in the path %s\n", r.ResponseFile, assetPath)

		return []byte{}
	}

	return data
}

//BuildStatusCode emulate an error code
func (r ResponseDefinition) BuildStatusCode() int {
	if r.Error == nil {
		return r.StatusCode
	}

	return r.StatusCode
}

//BuildHeaders build headres for the response
func (r ResponseDefinition) BuildHeaders() map[string]string {
	return r.Headers
}
