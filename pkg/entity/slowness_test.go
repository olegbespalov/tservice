package entity_test

import (
	"testing"
	"time"

	"github.com/olegbespalov/tservice/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestSlownessHappened(t *testing.T) {
	s := entity.Slowness{
		Chance: 0,
	}

	assert.False(t, s.Happened())

	s = entity.Slowness{
		Chance: 100,
	}

	assert.True(t, s.Happened())
}

func TestSlownessWait(t *testing.T) {
	s := entity.Slowness{}

	assert.Equal(t, 5*time.Second, s.Wait())

	s = entity.Slowness{
		Duration: "33ms",
	}

	assert.Equal(t, 33*time.Millisecond, s.Wait())
}
