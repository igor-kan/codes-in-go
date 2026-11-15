package mod_hyperloglog_cardinality

import "errors"

var (
	ErrHLLEmpty     = errors.New("hll: empty or nil input")
	ErrHLLBounds    = errors.New("hll: index or dimension out of bounds")
	ErrHLLInvalid   = errors.New("hll: invalid configuration or state")
	ErrHLLInvariant = errors.New("hll: structural invariant violated")
)
