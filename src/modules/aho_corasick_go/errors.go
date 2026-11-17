package mod_aho_corasick_go

import "errors"

var (
	ErrAhoCorasickInvalid   = errors.New("ahocorasick: invalid argument")
	ErrAhoCorasickBounds    = errors.New("ahocorasick: index out of bounds")
	ErrAhoCorasickEmpty     = errors.New("ahocorasick: empty container")
	ErrAhoCorasickInvariant = errors.New("ahocorasick: invariant violated")
)
