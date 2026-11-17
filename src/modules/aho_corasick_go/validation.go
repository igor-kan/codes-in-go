package mod_aho_corasick_go

func validatePatterns(patterns [][]byte) error {
	if len(patterns) == 0 {
		return ErrAhoCorasickEmpty
	}
	for _, p := range patterns {
		if len(p) == 0 {
			return ErrAhoCorasickInvalid
		}
	}
	return nil
}
