package config_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/olegbespalov/tservice/pkg/config"
	"github.com/olegbespalov/tservice/pkg/entity"
	"gopkg.in/go-playground/assert.v1"
)

func TestService(t *testing.T) {
	var data = []byte(`
default_definitions:
   response:
      status_code: 200
      response: '{"message":"default"}'
      headers:
         - Content-Type:application/json
   error:
      status_code: 500
      response: '{"error":"default"}'
      headers:
         - Content-Type:application/json   
responses:
   response1:
      path: /lorem/ipsum
      definition:
         status_code: 200
         response: '{"hello":"TService"}'
         headers:
            - Content-Type:application/json
            - x-version:123
   response2:
      path: /lorem
      definition:
         status_code: 200
         response_file: lorem.json
      slowness:
         chance: 30
         duration: 5s
   response3:
      path: /lorem/error
      definition:
         status_code: 200
         response_file: lorem.json
         headers:
            - Content-Type:application/json
      error:
         chance: 10
         definition:
            status_code: 500
            response: '{"error":":("}'
`)

	file, err := ioutil.TempFile("/tmp", "prefix")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	_, _ = file.Write(data)

	fileName := file.Name()

	cfg := config.NewService("8087", fileName, "/tmp")

	assert.Equal(t, 3, len(cfg.Rules()))

	rule, ok := cfg.Rules()["response1"]
	assert.Equal(t, true, ok)
	assert.Equal(t, entity.Rule{
		Path: "/lorem/ipsum",
		Definition: entity.Definition{
			StatusCode: 200,
			Response:   "{\"hello\":\"TService\"}",
			Headers: []string{
				"Content-Type:application/json",
				"x-version:123",
			},
		},
	}, rule)

	rule, ok = cfg.Rules()["response2"]
	assert.Equal(t, true, ok)
	assert.Equal(t, entity.Rule{
		Path: "/lorem",
		Definition: entity.Definition{
			StatusCode:   200,
			ResponseFile: "lorem.json",
		},
		Slowness: &entity.Slowness{
			Chance:   30,
			Duration: "5s",
		},
	}, rule)

	rule, ok = cfg.Rules()["response3"]
	assert.Equal(t, true, ok)
	assert.Equal(t, entity.Rule{
		Path: "/lorem/error",
		Definition: entity.Definition{
			StatusCode:   200,
			ResponseFile: "lorem.json",
			Headers: []string{
				"Content-Type:application/json",
			},
		},
		Error: &entity.Error{
			Chance: 10,
			Definition: entity.Definition{
				StatusCode: 500,
				Response:   "{\"error\":\":(\"}",
			},
		},
	}, rule)

	assert.Equal(t, entity.Definition{
		StatusCode: 200,
		Response:   "{\"message\":\"default\"}",
		Headers: []string{
			"Content-Type:application/json",
		},
	}, cfg.DefaultDefinition())

	assert.Equal(t, entity.Definition{
		StatusCode: 500,
		Response:   "{\"error\":\"default\"}",
		Headers: []string{
			"Content-Type:application/json",
		},
	}, cfg.DefaultError())
}

func TestConfigModify(t *testing.T) {
	file, err := ioutil.TempFile("/tmp", "prefix")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(`
responses:
   response1:
      path: /lorem/ipsum
      definition:
         status_code: 200
         response: '{"hello":"TService"}'
         headers:
            - Content-Type:application/json
            - x-version:123      
   `))

	assert.Equal(t, nil, err)

	fileName := file.Name()

	cfg := config.NewService("8087", fileName, "/tmp")

	assert.Equal(t, 1, len(cfg.Rules()))
	rule, ok := cfg.Rules()["response1"]
	assert.Equal(t, true, ok)
	assert.Equal(t, "/lorem/ipsum", rule.Path)
	assert.Equal(t, "{\"hello\":\"TService\"}", rule.Definition.Response)
	assert.Equal(t, 2, len(rule.Definition.Headers))

	_, err = file.Write([]byte(`
responses:
   response1:
      path: /lorem/ipsum
      definition:
         status_code: 200
         response: '{"hello":"changed"}'         
   `))

	assert.Equal(t, nil, err)

	time.Sleep(1 * time.Second)

	assert.Equal(t, 1, len(cfg.Rules()))
	rule, ok = cfg.Rules()["response1"]
	assert.Equal(t, true, ok)
	assert.Equal(t, "/lorem/ipsum", rule.Path)
	assert.Equal(t, "{\"hello\":\"changed\"}", rule.Definition.Response)
	assert.Equal(t, 0, len(rule.Definition.Headers))
}
