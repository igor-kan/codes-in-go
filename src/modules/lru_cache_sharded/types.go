// Package lru_cache_sharded provides core types and interfaces for Concurrent Sharded LRU Cache with TTL Eviction.
package lru_cache_sharded

import "errors"

const Alignment = 64

var (
	ErrNullPointer    = errors.New("lru_cache_sharded: null pointer or uninitialized context")
	ErrInvalidParam   = errors.New("lru_cache_sharded: invalid parameter provided")
	ErrOutOfBounds    = errors.New("lru_cache_sharded: index out of bounds")
	ErrAllocationFail = errors.New("lru_cache_sharded: allocation failed")
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
