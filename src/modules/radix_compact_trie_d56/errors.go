package mod_radix_compact_trie_d56

import "errors"

var (
	ErrTrieInvalid   = errors.New("radix_compact_trie_d56: invalid argument")
	ErrTrieBounds    = errors.New("radix_compact_trie_d56: index out of bounds")
	ErrTrieEmpty     = errors.New("radix_compact_trie_d56: empty input")
	ErrTrieInvariant = errors.New("radix_compact_trie_d56: invariant violated")
)
