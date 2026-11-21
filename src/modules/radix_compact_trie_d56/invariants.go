package mod_radix_compact_trie_d56

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrTrieInvariant
	}
	return nil
}
