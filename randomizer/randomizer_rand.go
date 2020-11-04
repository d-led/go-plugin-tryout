package randomizer

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/d-led/go-plugin-tryout/interfaces"
)

// Rand plugin state
type Rand struct {
	requests int64
}

// NewRand constructs a default Rand
func NewRand() *Rand {
	// an unsafe demo random seed
	rand.Seed(time.Now().UTC().UnixNano())
	return &Rand{
		requests: 0,
	}
}

// InjectRand constructs a default Rand
func InjectRand(randomizer *interfaces.Randomizer) {
	*randomizer = NewRand()
}

// Get returns a random int
func (g *Rand) Get() int {
	atomic.AddInt64(&g.requests, 1)
	return rand.Int()
}
