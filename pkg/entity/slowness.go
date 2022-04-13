package entity

import (
	"math/rand"
	"time"

	"github.com/olegbespalov/tservice/internal/parser"
)

// Slowness define if response will be slow
type Slowness struct {
	Chance   int
	Duration string
}

// Happened check if slowness happened
func (s Slowness) Happened() bool {
	return rand.Intn(100) <= s.Chance // nolint:gosec
}

// Wait how long should wait before response
func (s Slowness) Wait() time.Duration {
	wait, _ := parser.ParseInterval(5*time.Second, s.Duration)

	return wait
}
