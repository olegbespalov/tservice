package entity_test

import (
	"testing"

	"github.com/olegbespalov/tservice/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestHappened(t *testing.T) {
	e := entity.Error{
		Chance: 0,
	}

	assert.False(t, e.Happened())

	e = entity.Error{
		Chance: 100,
	}

	assert.True(t, e.Happened())
}
