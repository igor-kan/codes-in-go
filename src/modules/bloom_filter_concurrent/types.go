// Package bloom_filter_concurrent provides core types and interfaces for Atomic CAS-Synchronized Concurrent Bloom Filter.
package bloom_filter_concurrent

import "errors"

const Alignment = 64

var (
	ErrNullPointer    = errors.New("bloom_filter_concurrent: null pointer or uninitialized context")
	ErrInvalidParam   = errors.New("bloom_filter_concurrent: invalid parameter provided")
	ErrOutOfBounds    = errors.New("bloom_filter_concurrent: index out of bounds")
	ErrAllocationFail = errors.New("bloom_filter_concurrent: allocation failed")
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
