package mod_bipartite_kuhn_match_d56

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrKuhnInvalid
	}
	return nil
}
