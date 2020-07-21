package entity

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

//ResponseRules response rule
type ResponseRules struct {
	Method string
	Path   string

	Definition ResponseDefinition

	Error    *Error
	Slowness *Slowness
}

func (r ResponseRules) String() string {
	return fmt.Sprintf("%s -  %s\nHeaders: %v", r.Method, r.Path, r.Definition.Headers)
}

//ResponseDefinition response definition
type ResponseDefinition struct {
	StatusCode   int `yaml:"status_code"`
	Response     string
	ResponseFile string `yaml:"response_file"`
	Body         []byte
	Headers      []string `yaml:"headers,flow"`
}

//Error define if response will be error
type Error struct {
	Chance int

	Definition ResponseDefinition
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
func (r ResponseRules) Fit(method, path string) bool {
	return r.Path == path && (len(r.Method) == 0 || r.Method == method)
}

// BuildBody define what will be in the body
func (r *ResponseRules) BuildBody(assetPath string) []byte {
	if len(r.Definition.Response) > 0 {
		return []byte(r.Definition.Response)
	}

	if len(r.Definition.ResponseFile) == 0 {
		return []byte{}
	}

	data, err := ioutil.ReadFile(filepath.Clean(assetPath + string(os.PathSeparator) + r.Definition.ResponseFile))
	if err != nil {
		log.Printf("can't find a response file %s in the path %s\n", r.Definition.ResponseFile, assetPath)

		return []byte{}
	}

	return data
}

//BuildStatusCode emulate an error code
func (r ResponseRules) BuildStatusCode() int {
	if r.Error != nil {
		return r.Error.Definition.StatusCode
	}

	return r.Definition.StatusCode
}

//BuildHeaders build headres for the response
func (r ResponseRules) BuildHeaders() map[string]string {
	if len(r.Definition.Headers) == 0 {
		return map[string]string{}
	}

	headers := make(map[string]string, len(r.Definition.Headers))
	for _, v := range r.Definition.Headers {
		parts := strings.Split(v, ":")

		// TODO: make properly
		headers[parts[0]] = parts[1]
	}

	return headers
}
