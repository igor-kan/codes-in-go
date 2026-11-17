package mod_interval_coalesce_go

func verifyStrictDisjoint(spans []Span) error {
	for i := 1; i < len(spans); i++ {
		if spans[i-1].End+1 >= spans[i].Start {
			return ErrIntervalInvariant
		}
	}
	return nil
}
