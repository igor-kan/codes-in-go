// Package token_bucket_limiter provides core types and interfaces for Nanosecond-Resolution Token Bucket Rate Limiter.
package token_bucket_limiter

import "errors"

const Alignment = 64

var (
	ErrNullPointer    = errors.New("token_bucket_limiter: null pointer or uninitialized context")
	ErrInvalidParam   = errors.New("token_bucket_limiter: invalid parameter provided")
	ErrOutOfBounds    = errors.New("token_bucket_limiter: index out of bounds")
	ErrAllocationFail = errors.New("token_bucket_limiter: allocation failed")
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
