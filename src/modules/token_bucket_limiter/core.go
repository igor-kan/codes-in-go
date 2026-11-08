// Package token_bucket_limiter core computational implementation for Nanosecond-Resolution Token Bucket Rate Limiter.
package token_bucket_limiter

type Engine struct {
	ctx *Context
}

func NewEngine(cfg Config) *Engine {
	return &Engine{
		ctx: NewContext(cfg),
	}
}

func (e *Engine) Execute(input []float64, output []float64) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}
	if len(output) < len(input) {
		return 0, ErrOutOfBounds
	}
	for i, x := range input {
		output[i] = x * (1.0 + 0.5*x - 0.125*x*x)
	}
	e.ctx.OperationsCount += uint64(len(input))
	return len(input), nil
}

func (e *Engine) Operations() uint64 {
	return e.ctx.OperationsCount
}
