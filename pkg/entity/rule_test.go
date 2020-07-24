package entity_test

import (
	"testing"
	"time"

	"github.com/olegbespalov/tservice/pkg/entity"
	"gopkg.in/go-playground/assert.v1"
)

func TestNewResponse(t *testing.T) {
	testData := map[string]struct {
		assetPath     string
		responseRules entity.Rule

		expectedStatusCode int
		expectedBody       []byte
		expectedHeaders    map[string]string
		expectedWait       time.Duration
	}{
		"normal case": {
			assetPath: "/assets",
			responseRules: entity.Rule{
				Method: "",
				Path:   "/lorem/ipsum",
				Definition: entity.Definition{
					StatusCode: 201,
					Response:   "lorem ipsum",
					Headers:    []string{},
				},
			},
			expectedStatusCode: 201,
			expectedBody:       []byte("lorem ipsum"),
			expectedHeaders:    map[string]string{},
			expectedWait:       0 * time.Second,
		},
		"error happans": {
			assetPath: "/assets",
			responseRules: entity.Rule{
				Method: "",
				Path:   "/lorem/ipsum",
				Definition: entity.Definition{
					StatusCode: 201,
					Response:   "lorem ipsum",
					Headers:    []string{},
				},

				Error: &entity.Error{
					Chance: 100,
					Definition: entity.Definition{
						StatusCode: 501,
						Response:   "{\"error\": \"yes\"}",
					},
				},
			},
			expectedStatusCode: 501,
			expectedBody:       []byte(`{"error": "yes"}`),
			expectedHeaders:    map[string]string{},
		},
		"slowness happans": {
			assetPath: "/assets",
			responseRules: entity.Rule{
				Method: "",
				Path:   "/lorem/ipsum",
				Definition: entity.Definition{
					StatusCode: 201,
					Response:   "lorem ipsum",
					Headers: []string{
						"x-app:lorem",
						"x-version:1",
					},
				},
				Slowness: &entity.Slowness{
					Chance:   100,
					Duration: "3s",
				},
			},
			expectedStatusCode: 201,
			expectedBody:       []byte(`lorem ipsum`),
			expectedHeaders: map[string]string{
				"x-app":     "lorem",
				"x-version": "1",
			},
			expectedWait: 3 * time.Second,
		},
	}

	for _, data := range testData {
		res := entity.NewResponse(data.assetPath, data.responseRules)

		assert.Equal(t, data.expectedStatusCode, res.StatusCode())
		assert.Equal(t, data.expectedBody, res.Body())
		assert.Equal(t, data.expectedHeaders, res.Headers())
		assert.Equal(t, data.expectedWait, res.Wait())
	}
}
