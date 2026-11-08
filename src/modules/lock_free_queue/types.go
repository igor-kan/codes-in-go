// Package lock_free_queue provides core types and interfaces for Michael-Scott Non-Blocking Lock-Free Atomic Queue.
package lock_free_queue

import "errors"

const Alignment = 64

var (
	ErrNullPointer    = errors.New("lock_free_queue: null pointer or uninitialized context")
	ErrInvalidParam   = errors.New("lock_free_queue: invalid parameter provided")
	ErrOutOfBounds    = errors.New("lock_free_queue: index out of bounds")
	ErrAllocationFail = errors.New("lock_free_queue: allocation failed")
)

type Config struct {
	Capacity      int
	Tolerance     float64
	MaxIterations int
	Flags         uint32
}

func DefaultConfig() Config {
	return Config{
		Capacity:      1024,
		Tolerance:     1e-12,
		MaxIterations: 1000,
		Flags:         0,
	}
}

type Context struct {
	Config          Config
	OperationsCount uint64
	Buffer          []float64
}

func NewContext(cfg Config) *Context {
	return &Context{
		Config:          cfg,
		OperationsCount: 0,
		Buffer:          make([]float64, 0, cfg.Capacity),
	}
}
