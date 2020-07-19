package entity_test

import (
	"testing"
	"time"

	"github.com/olegbespalov/tservice/pkg/entity"
	"gopkg.in/go-playground/assert.v1"
)

func TestNewResponse(t *testing.T) {
	testData := map[string]struct {
		assetPath          string
		responseDefinition entity.ResponseDefinition

		expectedStatusCode int
		expectedBody       []byte
		expectedHeaders    map[string]string
		expectedWait       time.Duration
	}{
		"normal case": {
			assetPath: "/assets",
			responseDefinition: entity.ResponseDefinition{
				Method:     "",
				Path:       "/lorem/ipsum",
				StatusCode: 201,
				Response:   "lorem ipsum",
				Headers:    map[string]string{},
			},
			expectedStatusCode: 201,
			expectedBody:       []byte("lorem ipsum"),
			expectedHeaders:    map[string]string{},
		},
		"error happans": {
			assetPath: "/assets",
			responseDefinition: entity.ResponseDefinition{
				Method:     "",
				Path:       "/lorem/ipsum",
				StatusCode: 201,
				Response:   "lorem ipsum",
				Headers:    map[string]string{},
				Error: &entity.Error{
					Chance:     100,
					StatusCode: 501,
				},
			},
			expectedStatusCode: 501,
			expectedBody:       []byte(`{"error": "yes"}`),
			expectedHeaders:    nil,
		},
		"slowness happans": {
			assetPath: "/assets",
			responseDefinition: entity.ResponseDefinition{
				Method:     "",
				Path:       "/lorem/ipsum",
				StatusCode: 201,
				Response:   "lorem ipsum",
				Headers:    nil,
				Slowness: &entity.Slowness{
					Chance:   100,
					Duration: "3s",
				},
			},
			expectedStatusCode: 201,
			expectedBody:       []byte(`lorem ipsum`),
			expectedHeaders:    nil,
			expectedWait:       3 * time.Second,
		},
	}

	for _, data := range testData {
		res := entity.NewResponse(data.assetPath, data.responseDefinition)

		assert.Equal(t, data.expectedStatusCode, res.StatusCode())
		assert.Equal(t, data.expectedBody, res.Body())
		assert.Equal(t, data.expectedHeaders, res.Headers())
		assert.Equal(t, data.expectedWait, res.Wait())
	}
}
