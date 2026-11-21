package mod_radix_compact_trie_d56

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrTrieInvalid
	}
	return nil
}
