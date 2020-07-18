package parser_test

import (
	"testing"
	"time"

	"github.com/olegbespalov/tservice/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseInterval(t *testing.T) {

	testData := map[string]struct {
		defaultDuration  time.Duration
		input            string
		expectedDuration time.Duration
		expectedError    bool
	}{
		"2 seconds case": {
			defaultDuration:  13 * time.Second,
			input:            "2s",
			expectedDuration: 2 * time.Second,
			expectedError:    false,
		},
		"10 minuts case": {
			defaultDuration:  1 * time.Second,
			input:            "10m",
			expectedDuration: 10 * time.Minute,
			expectedError:    false,
		},
		"300 ms case": {
			defaultDuration:  13 * time.Second,
			input:            "300ms",
			expectedDuration: 300 * time.Millisecond,
			expectedError:    false,
		},
		"4 hours case": {
			defaultDuration:  13 * time.Second,
			input:            "4h",
			expectedDuration: 4 * time.Hour,
			expectedError:    false,
		},
		"default input": {
			defaultDuration:  77 * time.Second,
			input:            "",
			expectedDuration: 77 * time.Second,
			expectedError:    false,
		},
		"wrong input - duration": {
			defaultDuration:  1 * time.Second,
			input:            "1w",
			expectedDuration: 1 * time.Second,
			expectedError:    true,
		},
		"wrong input - format": {
			defaultDuration:  1 * time.Second,
			input:            "mm",
			expectedDuration: 1 * time.Second,
			expectedError:    true,
		},
	}

	for _, data := range testData {
		res, err := parser.ParseInterval(data.defaultDuration, data.input)

		assert.Equal(t, data.expectedDuration, res)
		assert.Equal(t, data.expectedError, err != nil)
	}
}
