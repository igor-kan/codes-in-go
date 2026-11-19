package mod_knuth_morris_pratt_go

import "errors"

var (
	ErrKmpGoInvalid   = errors.New("knuth_morris_pratt_go: invalid argument")
	ErrKmpGoBounds    = errors.New("knuth_morris_pratt_go: index out of bounds")
	ErrKmpGoEmpty     = errors.New("knuth_morris_pratt_go: empty input")
	ErrKmpGoInvariant = errors.New("knuth_morris_pratt_go: invariant violated")
)
