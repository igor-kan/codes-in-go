package mod_trie_wildcard_d59

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrWtrieInvalid
	}
	return nil
}
