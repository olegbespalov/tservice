package entity_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/olegbespalov/tservice/pkg/entity"
	"gopkg.in/go-playground/assert.v1"
)

func TestBuildBodyFromResponse(t *testing.T) {
	d := entity.Definition{
		Response:     "hey",
		ResponseFile: "lorem.fix",
	}

	assert.Equal(t, []byte("hey"), d.BuildBody("path"))
}

func TestBuildBodyFromFile(t *testing.T) {
	file, err := ioutil.TempFile("/tmp", "prefix")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	_, _ = file.WriteString("lorem content")

	s, _ := file.Stat()

	d := entity.Definition{
		ResponseFile: s.Name(),
	}

	assert.Equal(t, []byte("lorem content"), d.BuildBody("/tmp"))
}

func TestBuildBodyNoFile(t *testing.T) {
	d := entity.Definition{
		ResponseFile: "some-not-existing-file",
	}

	assert.Equal(t, []byte{}, d.BuildBody("/tmp"))
}

func TestBuildStatus(t *testing.T) {
	d := entity.Definition{
		StatusCode: http.StatusOK,
	}

	assert.Equal(t, http.StatusOK, d.BuildStatusCode())

	d = entity.Definition{
		StatusCode: 700,
	}

	assert.Equal(t, 700, d.BuildStatusCode())
}

func TestBuildHeaders(t *testing.T) {
	d := entity.Definition{}

	assert.Equal(t, map[string]string{}, d.BuildHeaders())

	d = entity.Definition{
		Headers: []string{
			"x-app:123",
			"x-version:3",
		},
	}

	assert.Equal(t, map[string]string{
		"x-app":     "123",
		"x-version": "3",
	}, d.BuildHeaders())
}
