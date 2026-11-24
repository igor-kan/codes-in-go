package mod_trie_wildcard_d59

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrWtrieInvariant
	}
	return nil
}
