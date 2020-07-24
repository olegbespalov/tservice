package entity

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//Definition response definition
type Definition struct {
	StatusCode   int `yaml:"status_code"`
	Response     string
	ResponseFile string `yaml:"response_file"`
	Body         []byte
	Headers      []string `yaml:"headers,flow"`
}

// BuildBody define what will be in the body
func (d Definition) BuildBody(assetPath string) []byte {
	if len(d.Response) > 0 {
		return []byte(d.Response)
	}

	if len(d.ResponseFile) == 0 {
		return []byte{}
	}

	data, err := ioutil.ReadFile(filepath.Clean(assetPath + string(os.PathSeparator) + d.ResponseFile))
	if err != nil {
		log.Printf("can't find a response file %s in the path %s\n", d.ResponseFile, assetPath)

		return []byte{}
	}

	return data
}

//BuildStatusCode emulate an error code
func (d Definition) BuildStatusCode() int {
	return d.StatusCode
}

//BuildHeaders build headres for the response
func (d Definition) BuildHeaders() map[string]string {
	if len(d.Headers) == 0 {
		return map[string]string{}
	}

	headers := make(map[string]string, len(d.Headers))
	for _, v := range d.Headers {
		parts := strings.Split(v, ":")

		// TODO: make properly
		headers[parts[0]] = parts[1]
	}

	return headers
}
