package pushmq

import (
	"github.com/panjf2000/ants/v2"
	"time"
)

const (
	// DefaultAntsPoolSize sets up the capacity of worker pool, 256 * 1024.
	DefaultAntsPoolSize = 1 << 8

	// ExpiryDuration is the interval time to clean up those expired workers.
	ExpiryDuration = 10 * time.Second

	// Nonblocking decides what to do when submitting a new task to a full worker pool: waiting for a available worker
	// or returning nil directly.
	Nonblocking = true
)

type Pool = ants.Pool

func DefaultPool() *Pool {
	options := ants.Options{ExpiryDuration: ExpiryDuration, Nonblocking: Nonblocking}
	defaultAntsPool, _ := ants.NewPool(DefaultAntsPoolSize, ants.WithOptions(options))
	return defaultAntsPool
}
