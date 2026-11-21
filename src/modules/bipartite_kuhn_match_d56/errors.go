package mod_bipartite_kuhn_match_d56

import "errors"

var (
	ErrKuhnInvalid   = errors.New("bipartite_kuhn_match_d56: invalid argument")
	ErrKuhnBounds    = errors.New("bipartite_kuhn_match_d56: index out of bounds")
	ErrKuhnEmpty     = errors.New("bipartite_kuhn_match_d56: empty input")
	ErrKuhnInvariant = errors.New("bipartite_kuhn_match_d56: invariant violated")
)
