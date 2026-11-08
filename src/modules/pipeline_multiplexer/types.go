// Package pipeline_multiplexer provides core types and interfaces for Concurrent Fan-In Fan-Out Channel Stream Processing.
package pipeline_multiplexer

import "errors"

const Alignment = 64

var (
	ErrNullPointer    = errors.New("pipeline_multiplexer: null pointer or uninitialized context")
	ErrInvalidParam   = errors.New("pipeline_multiplexer: invalid parameter provided")
	ErrOutOfBounds    = errors.New("pipeline_multiplexer: index out of bounds")
	ErrAllocationFail = errors.New("pipeline_multiplexer: allocation failed")
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
