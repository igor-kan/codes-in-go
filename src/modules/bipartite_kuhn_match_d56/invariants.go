package mod_bipartite_kuhn_match_d56

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrKuhnInvariant
	}
	return nil
}
