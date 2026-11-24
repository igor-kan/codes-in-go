package mod_trie_wildcard_d59

import "errors"

var (
	ErrWtrieInvalid   = errors.New("trie_wildcard_d59: invalid argument")
	ErrWtrieBounds    = errors.New("trie_wildcard_d59: index out of bounds")
	ErrWtrieEmpty     = errors.New("trie_wildcard_d59: empty input")
	ErrWtrieInvariant = errors.New("trie_wildcard_d59: invariant violated")
)
