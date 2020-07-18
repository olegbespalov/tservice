package entity

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Config for the service
type Config struct {
	Responses map[string]Response
}

// Load data from assest
func (c *Config) Load(assetPath string) {
	for k, v := range c.Responses {
		v.Body = body(assetPath, v)
		c.Responses[k] = v
	}
}

func body(assetPath string, r Response) []byte {
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
